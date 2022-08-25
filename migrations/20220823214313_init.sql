-- +goose Up
-- +goose StatementBegin
create table url_info
(
    url_info_id bigserial not null,
    user_id     bigint,
    short       text not null,
    url         text not null,
    created_at  timestamp without time zone not null default current_timestamp
);

create index idx_url_info_id on url_info(url_info_id);
create index idx_url_info_short on url_info(short);

create table follow
(
    url_info_id bigint not null,
    created_at  timestamp without time zone not null default current_timestamp
);

create index idx_follow on url_info(url_info_id);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table url_info;
drop table follow;
-- +goose StatementEnd
