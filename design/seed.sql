-- Enable pgcrypto for UUID generation
CREATE EXTENSION IF NOT EXISTS "pgcrypto";

-- Use core schema
SET search_path TO core;

-- ------------------------
-- Categories
-- ------------------------
INSERT INTO category (id, name, description)
VALUES
    (gen_random_uuid(), 'Electronics', 'Phones, Laptops, Gadgets'),
    (gen_random_uuid(), 'Books', 'Fiction, Non-fiction, Education'),
    (gen_random_uuid(), 'Home Appliances', 'Kitchen & Home Essentials'),
    (gen_random_uuid(), 'Fashion', 'Men & Women Clothing'),
    (gen_random_uuid(), 'Sports', 'Sporting Goods and Accessories'),
    (gen_random_uuid(), 'Beauty', 'Cosmetics and Skincare'),
    (gen_random_uuid(), 'Gaming', 'Consoles, Games, Accessories'),
    (gen_random_uuid(), 'Toys', 'Children Toys and Games'),
    (gen_random_uuid(), 'Health', 'Supplements & Equipment'),
    (gen_random_uuid(), 'Music', 'Instruments & Equipment');

-- ------------------------
-- Products
-- ------------------------
INSERT INTO product (
    id, category_id, name, description, redemption_points,
    stock_quantity, is_offer, tags, created_at
)
SELECT
    gen_random_uuid(),
    id,
    p.name,
    p.description,
    p.redemption_points,
    p.stock_quantity,
    p.is_offer,
    to_jsonb(string_to_array(p.tags, ',')),  -- convert to jsonb array
    NOW()
FROM (
         VALUES
             ('iPhone 15', 'Latest Apple smartphone', 1200, 30, true, 'electronics,apple,smartphone'),
             ('Samsung Galaxy S23', 'Flagship Android phone', 1100, 25, false, 'electronics,android,samsung'),
             ('Air Fryer', 'Oil-free cooking', 350, 40, true, 'home,appliance,kitchen'),
             ('Nike Sneakers', 'Men sport shoes', 500, 60, false, 'fashion,men,shoes'),
             ('Resistance Bands Set', 'Fitness kit', 200, 50, false, 'sports,fitness'),
             ('Acoustic Guitar', 'Beginner-friendly instrument', 650, 15, false, 'music,guitar,acoustic'),
             ('LEGO Super Mario', 'Interactive building set', 300, 20, true, 'toys,lego'),
             ('MacBook Air M2', 'Lightweight Apple laptop', 2200, 10, false, 'electronics,apple,laptop'),
             ('Cosrx Snail Cream', 'Korean skincare moisturizer', 150, 100, true, 'beauty,skincare'),
             ('Football Ball', 'FIFA-approved', 180, 45, false, 'sports,ball'),
             ('Hair Dryer', 'Powerful styling tool', 230, 35, false, 'beauty,hair'),
             ('Gaming Mouse', 'RGB wired mouse', 250, 40, true, 'gaming,mouse,accessories'),
             ('Digital Book Reader', 'Read anywhere', 480, 30, false, 'books,ebook,reader'),
             ('Protein Powder', 'Whey protein supplement', 320, 70, true, 'health,supplement'),
             ('T-shirt Pack', '3-Pack cotton tees', 100, 150, false, 'fashion,unisex,casual')
     ) AS p(name, description, redemption_points, stock_quantity, is_offer, tags),
     category;

-- ------------------------
-- Credit Packages
-- ------------------------
INSERT INTO credit_package (id, name, price_egp, reward_points, credits, is_active, created_at)
VALUES
    (gen_random_uuid(), 'Starter Pack', 150.00, 150, 150, true, NOW()),        -- Enough for basic redemptions
    (gen_random_uuid(), 'Pro Pack', 450.00, 500, 500, true, NOW()),            -- Mid-tier like Guitar, Air Fryer
    (gen_random_uuid(), 'Elite Pack', 900.00, 1000, 1000, true, NOW()),        -- Covers most products
    (gen_random_uuid(), 'Mega Pack', 1800.00, 2000, 2000, true, NOW()),        -- Can redeem MacBook, iPhone
    (gen_random_uuid(), 'Promo Pack', 19.99, 50, 50, true, NOW());

-- ------------------------
-- Users
-- ------------------------
INSERT INTO "user" (id, first_name, last_name, username, email, password_hash, role, status, created_at)
VALUES
    (gen_random_uuid(), 'Alice', 'Johnson', 'alicej', 'alice@example.com', 'hashedpwd1', 'user', 'active', NOW()),
    (gen_random_uuid(), 'Bob', 'Smith', 'bobsmith', 'bob@example.com', 'hashedpwd2', 'user', 'active', NOW()),
    (gen_random_uuid(), 'Charlie', 'Davis', 'charlied', 'charlie@example.com', 'hashedpwd3', 'admin', 'active', NOW()),
    (gen_random_uuid(), 'Diana', 'Prince', 'diana', 'diana@example.com', 'hashedpwd4', 'user', 'active', NOW()),
    (gen_random_uuid(), 'Ethan', 'Hunt', 'ethanh', 'ethan@example.com', 'hashedpwd5', 'user', 'suspended', NOW()),
    (gen_random_uuid(), 'Fiona', 'Lee', 'flee', 'fiona@example.com', 'hashedpwd6', 'user', 'active', NOW()),
    (gen_random_uuid(), 'George', 'Hall', 'georgeh', 'george@example.com', 'hashedpwd7', 'user', 'banned', NOW()),
    (gen_random_uuid(), 'Hana', 'Kim', 'hana_k', 'hana@example.com', 'hashedpwd8', 'user', 'active', NOW()),
    (gen_random_uuid(), 'Ivan', 'Petrov', 'ivanp', 'ivan@example.com', 'hashedpwd9', 'user', 'active', NOW()),
    (gen_random_uuid(), 'Jane', 'Doe', 'janed', 'jane@example.com', 'hashedpwd10', 'user', 'active', NOW());

-- ------------------------
-- Wallets (linked to users above)
-- ------------------------
INSERT INTO wallet (id, user_id, credits_balance, points_balance, updated_at)
SELECT gen_random_uuid(), id, (random() * 1000)::int, (random() * 2000)::int, NOW()
FROM "user";

-- ------------------------
-- Purchases (linked to credit packages + users)
-- ------------------------
INSERT INTO purchase (id, user_id, credit_package_id, status, credits, created_at)
SELECT gen_random_uuid(), u.id, cp.id,
       (ARRAY['completed','pending'])[floor(random() * 2 + 1)],
       cp.credits, NOW() - (random() * INTERVAL '30 days')
FROM "user" u
    JOIN credit_package cp ON true
    LIMIT 20;

-- ------------------------
-- Redemptions (linked to users + products)
-- ------------------------
INSERT INTO redemption (id, user_id, product_id, quantity, status, created_at)
SELECT gen_random_uuid(), u.id, p.id,
       (1 + floor(random() * 3))::int,
    (ARRAY['pending','delivered','cancelled'])[floor(random() * 3 + 1)],
       NOW() - (random() * INTERVAL '15 days')
FROM "user" u
    JOIN product p ON true
    LIMIT 20;
