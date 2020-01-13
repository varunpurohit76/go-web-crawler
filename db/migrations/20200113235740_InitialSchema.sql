-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE `url`
(
    `data` varchar(512) NOT NULL DEFAULT '',
    `id`   varchar(512) NOT NULL DEFAULT ''
);
CREATE TABLE `relation`
(
    `child`  varchar(512) NOT NULL DEFAULT '',
    `parent` varchar(512) NOT NULL DEFAULT ''
);
-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE `url`;
DROP TABLE `relation`;