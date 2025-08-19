

CREATE TABLE IF NOT EXISTS orders {
    id CHAR(27) PRIMARY KEY,
    created_at TIMESTAMP WITH TIME ZONE NOT NULL,
    account_id CHAR(24) NOT NULL,
    total_price MONEY NOT NULL,
}

