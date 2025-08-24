-- Remove as colunas adicionadas
ALTER TABLE categories 
DROP COLUMN updated_at,
DROP COLUMN is_active;
