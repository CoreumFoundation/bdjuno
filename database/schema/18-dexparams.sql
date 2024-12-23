/* ---- PARAMS ---- */

CREATE TABLE dex_params
(
    one_row_id BOOLEAN NOT NULL DEFAULT TRUE PRIMARY KEY,
    params     JSONB   NOT NULL,
    height     BIGINT  NOT NULL,
    CHECK (one_row_id)
);
CREATE INDEX dex_params_height_index ON dex_params (height);
