# jwkserver
使用gin框架和postgresql数据库搭建的jwkserver

## 建库脚本
-- SCHEMA: go

-- DROP SCHEMA go ;

CREATE SCHEMA go
    AUTHORIZATION pg;

-- Table: go.jwkeys

-- DROP TABLE go.jwkeys;

CREATE TABLE go.jwkeys
(
    ulid character varying(100) COLLATE pg_catalog."default" NOT NULL,
    privatekey json NOT NULL,
    publickey json NOT NULL,
    creatstime text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT jwks_pkey PRIMARY KEY (ulid)
)
WITH (
    OIDS = FALSE
)
TABLESPACE pg_default;

ALTER TABLE go.jwkeys
    OWNER to pg;

## 初始化运行命令

1、git clone https://github.com/qrcode-star/jwkserver.git
2、CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o jwkserver jwkserver.go
3、docker build -t jwkserver .
4、docker run --name jwkserver -p 8080:80 -d jwkserver -p=172.18.24.82
