-- Adiciona a coluna updated_at à tabela categories
ALTER TABLE categories 
ADD COLUMN updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
ADD COLUMN is_active BOOLEAN DEFAULT true;
