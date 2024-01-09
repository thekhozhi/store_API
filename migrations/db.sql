CREATE TABLE users (
    id uuid primary key,
    first_name varchar(30),
    last_name varchar(30),
    email varchar(50),
    phone varchar(9)
);

CREATE TABLE products (
    id uuid primary key,
    name varchar(30),
    price int
);

CREATE TABLE orders (
    id int primary key,
    amount int,
    user_id uuid references users(id),
    Created_at varchar(10)
);

CREATE TABLE order_products(
    id uuid,
    order_id int references orders(id),
    product_id uuid references products(id),
    quantity int,
    price int
);

INSERT INTO users values (
    'f3d55d0c-4213-41a7-a3fa-380b1f53e170',
    'Asadbek',
    'Baxodirov',
    'asadbaxodir@gmail.com',
    '903214789'
);

INSERT INTO products values(
    'bd5e0a1c-c37b-405a-8ec8-3430746c86a3',
    'suv',
    3000
);

INSERT INTO orders (id, amount, created_at) values(
    1,
    15000,
    '10-01-2024'
);

INSERT INTO order_products (id, quantity, price) values(
    'c180dc8b-c59a-44b0-a0f3-3d57a981ef3c',
    5,
    3000
);