CREATE TABLE account_denom_balance
(
    account         TEXT        NOT NULL,
    denom           TEXT        NOT NULL,
    amount          TEXT        NOT NULL,
    PRIMARY KEY(account, denom)
);

CREATE VIEW token_holder_count AS
    SELECT denom, COUNT(*) AS holders
    FROM account_denom_balance
    GROUP BY denom;
