DO $$
BEGIN
    IF NOT EXISTS (
        SELECT FROM pg_tables
        WHERE schemaname = 'public' 
        AND tablename = 'projects'
    ) THEN
        CREATE TABLE projects (
            name TEXT PRIMARY KEY,
            display_name TEXT,
            description TEXT,
            color_tag TEXT,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            state INT
        );
        
        RAISE NOTICE 'Table created successfully';
    ELSE
        RAISE NOTICE 'Table already exists';
    END IF;
END $$;
