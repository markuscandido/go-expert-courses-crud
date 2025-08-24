-- Drop triggers
DROP TRIGGER IF EXISTS update_courses_updated_at ON courses;

-- Only drop the function if no other triggers are using it
DO $$
BEGIN
    IF NOT EXISTS (
        SELECT 1 FROM pg_trigger 
        WHERE tgname = 'update_categories_updated_at'
    ) THEN
        DROP FUNCTION IF EXISTS update_updated_at_column();
    END IF;
END $$;

-- Drop columns from courses
ALTER TABLE courses 
DROP COLUMN IF EXISTS updated_at,
DROP COLUMN IF EXISTS is_active;
