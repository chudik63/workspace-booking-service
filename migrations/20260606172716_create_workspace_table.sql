-- +goose Up
-- +goose StatementBegin
CREATE TABLE workspace (
    id           BIGSERIAL    PRIMARY KEY,
    name         TEXT         NOT NULL,
    full_address TEXT         NOT NULL,
    lat          NUMERIC(12, 5) NOT NULL,
    lon          NUMERIC(12, 5) NOT NULL,
    type         TEXT     NOT NULL,
    status       TEXT NOT NULL,
    capacity     INTEGER      NOT NULL,
    created_at   TIMESTAMP  NOT NULL DEFAULT now(),
    updated_at   TIMESTAMP  NOT NULL DEFAULT now(),
    is_deleted   BOOLEAN      NOT NULL DEFAULT FALSE
);

CREATE UNIQUE INDEX workspace_unique_name_type_idx ON workspace (name, type) WHERE is_deleted = false;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE workspace;
-- +goose StatementEnd
