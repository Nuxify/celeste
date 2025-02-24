CREATE TABLE
    `user_sss` (
        `user_wallet_address` varchar(42) NOT NULL,
        `sss_3` varchar(100) NOT NULL,
        `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
        PRIMARY KEY (`user_wallet_address`),
        FOREIGN KEY (`user_wallet_address`) REFERENCES users(`wallet_address`)
    ) ENGINE = InnoDB DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_0900_ai_ci;