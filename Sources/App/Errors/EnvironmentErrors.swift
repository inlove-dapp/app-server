//
//  File.swift
//
//
//  Created by Qiwei Li on 1/10/24.
//

import Foundation

enum EnvironmentErrors: LocalizedError {
    case MissingRequiredEnvironment(key: String)
    
    var errorDescription: String? {
        switch self {
        case .MissingRequiredEnvironment(let key):
            return "Missing required environment variable: \(key)"
        }
    }
}
