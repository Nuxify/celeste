CREATE TABLE
    `users` (
        `wallet_address` varchar(42) NOT NULL,
        `email` varchar(100) NOT NULL UNIQUE,
        `password` varchar(100) NOT NULL,
        `sss_1` varchar(100) NOT NULL,
        `name` varchar(255) NOT NULL,
        `email_verified_at` timestamp NULL DEFAULT NULL,
        `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
        `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
        PRIMARY KEY (`wallet_address`)
 );