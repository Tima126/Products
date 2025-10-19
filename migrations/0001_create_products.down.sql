CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL
);

INSERT INTO products (name, description, price) VALUES
('iPhone 15', 'Новый iPhone 15', 1200.00),
('Samsung Galaxy S23', 'Флагман Samsung', 1000.00),
('Google Pixel 8', 'Новый Pixel 8', 900.00),
('OnePlus 12', 'Флагман OnePlus', 800.00),
('Xiaomi 14', 'Флагман Xiaomi', 700.00);