//
//  SupabaseMiddleware.swift
//
//
//  Created by Qiwei Li on 1/10/24.
//

import Foundation
import JWTKit
import Vapor

struct Payload: JWTPayload {
    func verify(using signer: JWTKit.JWTSigner) throws {}
}

struct SupabaseMiddleware: AsyncMiddleware {
    func respond(to request: Request, chainingTo next: AsyncResponder) async throws -> Response {
        // get bearer token from request
        let bearerToken = request.headers.bearerAuthorization?.token
        guard let bearerToken = bearerToken else {
            throw Abort(.unauthorized, reason: "Missing bearer token")
        }
        let signer: JWTSigner = .hs256(key: Environments.SupabaseJwtToken)
        do {
            _ = try signer.verify(bearerToken, as: Payload.self)
        } catch {
            throw Abort(.unauthorized, reason: "Invalid bearer token")
        }
        return try await next.respond(to: request)
    }
}
