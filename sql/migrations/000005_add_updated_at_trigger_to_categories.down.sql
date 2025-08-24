-- Drop the trigger
DROP TRIGGER IF EXISTS update_categories_updated_at ON categories;

-- Drop the function
DROP FUNCTION IF EXISTS update_category_updated_at_column();
