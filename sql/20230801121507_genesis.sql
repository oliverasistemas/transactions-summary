-- Create the 'accounts' table
use stori;
CREATE TABLE accounts (
                          id INT PRIMARY KEY AUTO_INCREMENT,
                          account_name VARCHAR(100) NOT NULL
);

-- Create the 'transactions' table
CREATE TABLE transactions (
                              id INT PRIMARY KEY AUTO_INCREMENT,
                              transaction_date DATE NOT NULL,
                              transaction_amount DECIMAL(10, 2) NOT NULL,
                              transaction_type ENUM('debit', 'credit') NOT NULL,
                              account_id INT NOT NULL,
                              FOREIGN KEY (account_id) REFERENCES accounts(id)
);

INSERT INTO accounts(account_name) VALUES ("ACME");
INSERT INTO accounts(account_name) VALUES ("DEMO");