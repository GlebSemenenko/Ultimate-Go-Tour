-- Таблица 1: Клиенты
CREATE TABLE Clients (
    client_id INT PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    email VARCHAR(150),
    phone VARCHAR(20),
    city VARCHAR(100)
);

-- Таблица 2: Товары
CREATE TABLE Products (
    product_id INT PRIMARY KEY,
    name VARCHAR(150) NOT NULL,
    description VARCHAR(500),
    price DECIMAL(10, 2) NOT NULL,
    category VARCHAR(50)
);

-- Таблица 3: Заказы (связывает клиентов и товары)
CREATE TABLE Orders (
    order_id INT PRIMARY KEY,
    client_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL DEFAULT 1,
    order_date DATE NOT NULL,
    total_price DECIMAL(10, 2),
    FOREIGN KEY (client_id) REFERENCES Clients(client_id),
    FOREIGN KEY (product_id) REFERENCES Products(product_id)
);