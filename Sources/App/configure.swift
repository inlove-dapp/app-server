import Fluent
import FluentMongoDriver
import NIOSSL
import Vapor

// configures your application
public func configure(_ app: Application) async throws {
    try app.databases.use(.mongo(connectionString: Environments.DatabaseUrl), as: .mongo)

    app.middleware.use(SupabaseMiddleware())
    app.migrations.add(CreateItineraryReport())

    // register routes
    try routes(app)
}
