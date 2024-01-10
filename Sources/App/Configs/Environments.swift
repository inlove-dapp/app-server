//
//  File.swift
//
//
//  Created by Qiwei Li on 1/10/24.
//

import Foundation
import Vapor

enum Environments {
    /**
     Supabase JWT Signing Secret
     */
    static let SupabaseJwtToken = Environments.guardGetEnvironment("SUPABASE_JWT_SECRET")
    /**
     Database connection url
     */
    static let DatabaseUrl = Environments.guardGetEnvironment("DATABASE_URL")

    static func guardGetEnvironment(_ key: String) -> String {
        guard let value = Environment.get(key) else {
            fatalError(EnvironmentErrors.MissingRequiredEnvironment(key: key).errorDescription!)
        }

        return value
    }
}
