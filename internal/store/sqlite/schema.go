package sqlite

const (

	enableForeignKeys = `PRAGMA foreign_keys = ON;`

	createTableUsers = `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY,
		username TEXT NOT NULL UNIQUE,
		email TEXT NOT NULL UNIQUE,
		password_hash TEXT NOT NULL,
		last_login DATETIME DEFAULT CURRENT_TIMESTAMP,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		remembered_token TEXT,
		is_admin BOOLEAN NOT NULL
	);
	`
	createTableArtists = `
	CREATE TABLE IF NOT EXISTS artists (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT,
		spotify_id TEXT
	);
	`

	createTableGenres = `
	CREATE TABLE IF NOT EXISTS genres (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL
	);
	`

	createTableVenues = `
	CREATE TABLE IF NOT EXISTS venues (
		id INTEGER PRIMARY KEY,
		name TEXT NOT NULL,
		address TEXT NOT NULL,
		city TEXT NOT NULL,
		longitude REAL NOT NULL,
		latitude REAL NOT NULL
	);
	`

	createTableGigs = `
	CREATE TABLE IF NOT EXISTS gigs (
		id INTEGER PRIMARY KEY,
		venue_id INTEGER NOT NULL,
		date_time DATETIME NOT NULL,
		name TEXT NOT NULL,
		description TEXT,
		ticket_url TEXT,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (venue_id) REFERENCES venues(id) ON DELETE CASCADE
	);
	`


	// --- Join Tables --- //
	createTableGigArtists = `
	CREATE TABLE IF NOT EXISTS gig_artists (
		gig_id INTEGER NOT NULL,
		artist_id INTEGER NOT NULL,
		PRIMARY KEY (gig_id, artist_id),
		FOREIGN KEY (gig_id) REFERENCES gigs(id) ON DELETE CASCADE,
		FOREIGN KEY (artist_id) REFERENCES artists(id) ON DELETE CASCADE
	);
	`

	createTableGigGenres = `
	CREATE TABLE IF NOT EXISTS gig_genres (
		gig_id INTEGER NOT NULL,
		genre_id INTEGER NOT NULL,
		PRIMARY KEY (gig_id, genre_id),
		FOREIGN KEY (gig_id) REFERENCES gigs(id) ON DELETE CASCADE,
		FOREIGN KEY (genre_id) REFERENCES genres(id) ON DELETE CASCADE
	);
	`

	createTableArtistGenres = `
	CREATE TABLE IF NOT EXISTS artist_genres (
		artist_id INTEGER NOT NULL,
		genre_id INTEGER NOT NULL,
		PRIMARY KEY (artist_id, genre_id),
		FOREIGN KEY (artist_id) REFERENCES artists(id) ON DELETE CASCADE,
		FOREIGN KEY (genre_id) REFERENCES genres(id) ON DELETE CASCADE
	);
	`

	// --- Indexes --- //
	createIndexGigDateTime = `CREATE INDEX IF NOT EXISTS idx_gig_date_time ON gigs(date_time);`
	createIndexGigVenueID = `CREATE INDEX IF NOT EXISTS idx_gig_venue_id ON gigs(venue_id);`
	createIndexGigArtists = `CREATE INDEX IF NOT EXISTS idx_gig_artists ON gig_artists(artist_id);`

	createIndexGigGenresGigID = `CREATE INDEX IF NOT EXISTS idx_gig_genres_gig_id ON gig_genres(gig_id);`
	createIndexGigGenresGenreID = `CREATE INDEX IF NOT EXISTS idx_gig_genres_genre_id ON gig_genres(genre_id);`

	createIndexArtistName = `CREATE INDEX IF NOT EXISTS idx_artist_name ON artists(name);`
	createIndexVenueName = `CREATE INDEX IF NOT EXISTS idx_venue_name ON venues(name);`
	
	createIndexUserEmail = `CREATE INDEX IF NOT EXISTS idx_user_email ON users(email);`
)

// Schema statements grouped for execution
var schemaStatements = []string{
	enableForeignKeys,
	
	createTableUsers,
	createTableArtists,
	createTableGenres,
	createTableVenues,
	createTableGigs,
	createTableGigArtists,
	createTableGigGenres,
	createTableArtistGenres,

	createIndexGigDateTime,
	createIndexGigVenueID,
	createIndexGigArtists,
	createIndexGigGenresGigID,
	createIndexGigGenresGenreID,
	createIndexArtistName,
	createIndexVenueName,
	createIndexUserEmail,
}