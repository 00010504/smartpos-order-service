CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TYPE user_type AS ENUM (
    '1fe92aa8-2a61-4bf1-b907-182b497584ad', -- system user
    '9fb3ada6-a73b-4b81-9295-5c1605e54552'  -- admin user
);

CREATE TYPE app_type AS ENUM (
    '1fe92aa8-2a61-4bf1-b907-182b497584ad', -- client
    '9fb3ada6-a73b-4b81-9295-5c1605e54552'  -- admin
);

CREATE TYPE order_discount_type AS ENUM (
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',  -- none_type
    '1fe92aa8-2a61-4bf1-b907-182b497584ad', -- percentage
    '9fb3ada6-a73b-4b81-9295-5c1605e54552'  -- amount
);

CREATE TABLE IF NOT EXISTS "user" (
    "id" UUID PRIMARY KEY,
    "image" VARCHAR,
    "user_type_id" user_type NOT NULL,
    "first_name" VARCHAR(250) NOT NULL,
    "last_name" VARCHAR(250) NOT NULL,
    "phone_number" VARCHAR(30) NOT NULL,
    "deleted_at" BIGINT NOT NULL DEFAULT 0
);

CREATE INDEX "user_deleted_at_idx" ON "user"("deleted_at");

INSERT INTO "user" (
    "id",
    "first_name",
    "last_name",
    "phone_number",
    "user_type_id"
) VALUES (
    '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
    'admin',
    'admin',
    '99894172774',
    '9fb3ada6-a73b-4b81-9295-5c1605e54552'
);

CREATE TABLE "company" (
  "id" UUID PRIMARY KEY,
  "name" varchar(200) NOT NULL,
  "created_by" UUID,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" BIGINT NOT NULL DEFAULT 0,
  "deleted_by" UUID,
  UNIQUE ("name", "deleted_at")
);

CREATE INDEX "company_deleted_at_idx" ON "company"("deleted_at");


CREATE TABLE "shop" (
  "id" UUID PRIMARY KEY ,
  "title" VARCHAR NOT NULL,
  "company_id" UUID NOT NULL,
  "created_at" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  "created_by" UUID,
  "deleted_at" BIGINT NOT NULL DEFAULT 0,
  "deleted_by" UUID,
  UNIQUE ("title", "company_id", "deleted_at")
);
CREATE INDEX "shop_deleted_at_idx" ON "shop"("deleted_at");

CREATE TABLE "payment_type" (
  "id" UUID PRIMARY KEY,
  "name" varchar NOT NULL,
  "company_id" UUID NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_by" UUID,
  "deleted_at" BIGINT NOT NULL DEFAULT 0,
  "deleted_by" UUID,
  UNIQUE ("name", "company_id", "deleted_at")
);

CREATE INDEX "payment_type_deleted_at_idx" ON "payment_type"("deleted_at");

CREATE TABLE "cashbox" (
  "id" UUID PRIMARY KEY,
  "company_id" UUID ,
  "shop_id" UUID,
  "title" VARCHAR(200) NOT NULL,
  "cheque_id" UUID NOT NULL,
  "created_by" UUID,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" BIGINT NOT NULL DEFAULT 0,
  "deleted_by" UUID,
  UNIQUE ("company_id", "title", "deleted_at")
);
CREATE INDEX "cashbox_deleted_at_idx" ON "cashbox"("deleted_at");

CREATE TABLE "cheque" (
  "id" UUID PRIMARY KEY,
  "company_id" UUID NOT NULL,
  "name" VARCHAR(300) NOT NULL,
  "message" TEXT  NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" BIGINT NOT NULL DEFAULT 0,
  UNIQUE ("company_id", "name", "deleted_at")
);

CREATE INDEX "cheque_deleted_at_idx" ON "cheque"("deleted_at"); 


CREATE TABLE "receipt_block" (
  "id" UUID PRIMARY KEY,
  "name" VARCHAR(200) NOT NULL,
  "name_tr" JSONB NOT NULL,
  "deleted_at" BIGINT NOT NULL DEFAULT 0
);
CREATE INDEX "receipt_block_deleted_at_idx" ON "receipt_block"("deleted_at");

-- insert receipt block
INSERT INTO "receipt_block" ("id", "name", "name_tr") VALUES
    ('9bdde8d7-1b48-4788-9f6f-aa94bd6d9006', 'information_block', '{"en":"Information Block", "uz":"Axborot bloki", "ru":"Информационный блок"}'),
    ('a7655040-2942-4b61-9771-847f4f48a33d', 'bottom_block', '{"en":"Bottom Block", "uz":"Pastki blok", "ru":"Нижний блок"}');


CREATE TABLE "receipt_field"  (
  "id" UUID PRIMARY KEY,
  "name" VARCHAR(200) NOT NULL,
  "name_tr" JSONB NOT NULL,
  "block_id" UUID NOT NULL REFERENCES "receipt_block"("id") ON DELETE CASCADE,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" BIGINT NOT NULL DEFAULT 0,
  UNIQUE ("name", "deleted_at")
);
CREATE INDEX "receipt_field_deleted_at_idx" ON "receipt_field"("deleted_at");

INSERT INTO "receipt_field"("id", "name", "name_tr", "block_id" ) VALUES
    ('9527f6aa-6432-4932-ab43-5bfed278e751', 'shop name', '{"en":"Shop Name", "ru":"Название магазина", "uz":"Do''kon nomi"}', '9bdde8d7-1b48-4788-9f6f-aa94bd6d9006'),
    ('bc289940-a175-4460-a743-0aa3da9f8c7c', 'datetime', '{"en":"Datetime", "ru":"Дата и время", "uz":"Sana vaqti"}', '9bdde8d7-1b48-4788-9f6f-aa94bd6d9006'),
    ('1888a21b-110a-41a6-8c97-5d7b8f34ea3b', 'seller', '{"en":"Seller", "ru":"Продавец", "uz":"Sotuvchi"}', '9bdde8d7-1b48-4788-9f6f-aa94bd6d9006'),
    ('c2e9abae-1c25-4ac1-85c3-4c8c33d0d975', 'cashier', '{"en":"Cashier", "ru":"Кассир", "uz":"Kassir"}', '9bdde8d7-1b48-4788-9f6f-aa94bd6d9006'),
    ('15f8c6fe-8ef2-4eb2-bcac-1cf52402a4cd', 'customer', '{"en":"Customer", "ru":"Клиент", "uz":"mijoz"}', '9bdde8d7-1b48-4788-9f6f-aa94bd6d9006'),
    ('ab9c342d-8d6f-4e82-9622-11437c6e135c', 'contacts', '{"en":"Contacs", "ru":"Контакты", "uz":"Kontaktlar"}', '9bdde8d7-1b48-4788-9f6f-aa94bd6d9006');


CREATE  TABLE "cheque_logo" (
  "image" TEXT NOT NULL,
  "cheque_id" UUID NOT NULL,
  "left" INT NOT NULL DEFAULT 0,
  "right" INT NOT NULL DEFAULT 0,
  "top" INT NOT NULL DEFAULT 0,
  "bottom" INT NOT NULL DEFAULT 0,
  PRIMARY KEY("cheque_id")
);

CREATE TABLE "cheque_field" (
  "field_id" UUID NOT NULL,
  "cheque_id" UUID NOT NULL,
  "position" INT NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "deleted_at" BIGINT NOT NULL DEFAULT 0,
  PRIMARY KEY ("field_id", "cheque_id", "deleted_at")
);
CREATE INDEX "cheque_field_deleted_at" ON "cheque_field"("deleted_at");


CREATE TABLE IF NOT EXISTS "measurement_unit" (
    "id" UUID PRIMARY KEY,
    "company_id" UUID NOT NULL,
    "is_deletable" BOOLEAN NOT NULL DEFAULT TRUE,
    "short_name" VARCHAR NOT NULL,
    "long_name" VARCHAR NOT NULL,
    "precision" VARCHAR NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "created_by" UUID,
    "deleted_at" BIGINT NOT NULL DEFAULT 0,
    "deleted_by" UUID,
    UNIQUE ("short_name", "company_id", "deleted_at")
);
CREATE INDEX measurement_unit_deleted_at_idx ON "measurement_unit"("deleted_at");

CREATE TABLE "product" (
  "id" UUID PRIMARY KEY,
  "is_marking" BOOLEAN NOT NULL DEFAULT FALSE,
  "sku" VARCHAR NOT NULL,
  "measurement_unit_id" UUID NOT NULL,
  "mxik_code" VARCHAR,
  -- "brand_id" UUID,
  "image" VARCHAR,
  "name" varchar NOT NULL,
  "company_id" UUID NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_by" UUID,
  "deleted_at" BIGINT NOT NULL DEFAULT 0,
  "deleted_by" UUID,
  UNIQUE ("sku", "company_id", "deleted_at")
);
CREATE INDEX "product_deleted_at_idx" ON "product"("deleted_at");

CREATE TABLE IF NOT EXISTS "product_barcode" (
    "barcode" VARCHAR(300) NOT NULL,
    "product_id" UUID NOT NULL,
    PRIMARY KEY ("barcode", "product_id")
);

CREATE TABLE IF NOT EXISTS "shop_price" (
    "id" UUID,
    "supply_price" NUMERIC NOT NULL DEFAULT 0,
    "min_price" NUMERIC NOT NULL DEFAULT 0,
    "max_price" NUMERIC NOT NULL DEFAULT 0,
    "retail_price" NUMERIC NOT NULL DEFAULT 0,
    "whole_sale_price" NUMERIC NOT NULL DEFAULT 0,
    "shop_id" UUID NOT NULL,
    "product_id" UUID NOT NULL,
    PRIMARY KEY("product_id", "shop_id")
);

CREATE TABLE IF NOT EXISTS "measurement_values" (
    "shop_id" UUID NOT NULL,
    "is_available" BOOLEAN NOT NULL DEFAULT TRUE,
    "in_stock" NUMERIC NOT NULL DEFAULT 0,
    "product_id" UUID NOT NULL,
    PRIMARY KEY("product_id", "shop_id")
);

CREATE TABLE "client" (
  "id" UUID PRIMARY KEY,
  "first_name" VARCHAR NOT NULL,
  "last_name" VARCHAR NOT NULL,
  "phone_number" varchar NOT NULL,
  "company_id" UUID NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_by" UUID,
  "deleted_at" BIGINT NOT NULL DEFAULT 0,
  "deleted_by" UUID,
  UNIQUE ("phone_number", "company_id", "deleted_at")
);
CREATE INDEX "client_deleted_at_idx" ON "client"("deleted_at");

CREATE TABLE "order_status" (
  "id" UUID PRIMARY KEY,
  "name" varchar NOT NULL,
  "translation" JSONB NOT NULL,
  UNIQUE ("name")
);

INSERT INTO "order_status"(id, name, translation)
VALUES
('7069e210-7d2e-4a12-9160-3ef82f18ef4d', 'draft', '{"uz": "Qoralama", "ru": "Черновик", "en": "Draft"}'),
('d3bde6a2-532c-4f08-811f-0385e804c885', 'payed', '{"uz": "T''olangan", "ru": "Оплаченный", "en": "Payed"}'),
('15c0d291-45e5-4077-89d7-9b365e65cfed', 'postpone', '{"uz": "Kechiktirish", "ru": "Oткладывать", "en": "Postpone"}');

CREATE TABLE "order" (
  "id" UUID PRIMARY KEY,
  "external_id" VARCHAR NOT NULL,
  "status" UUID NOT NULL DEFAULT '7069e210-7d2e-4a12-9160-3ef82f18ef4d',
  "total_price" NUMERIC DEFAULT 0 NOT NULL,
  "total_discount_price" NUMERIC NOT NULL DEFAULT 0,
  "custom_discount_type" order_discount_type NOT NULL DEFAULT '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
  "custom_discount_value" NUMERIC NOT NULL DEFAULT 0,
  "product_sort_count" INT NOT NULL DEFAULT 0,
  "shop_id" UUID NOT NULL,
  "cashier_id" UUID NOT NULL,
  "seller_id" UUID,
  "client_id" UUID REFERENCES "client"("id") ON DELETE SET NULL,
  "cashbox_id" UUID NOT NULL,
  "company_id" UUID NOT NULL,
  "payed_time" TIMESTAMP,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_by" UUID,
  "deleted_at" BIGINT NOT NULL DEFAULT 0,
  "deleted_by" UUID,
  UNIQUE ("external_id", "shop_id", "deleted_at")
);
CREATE INDEX "order_deleted_at_idx" ON "order"("deleted_at");

CREATE TABLE "order_seller" (
  "user_id" UUID NOT NULL,
  "order_id" UUID NOT NULL,
  UNIQUE ("user_id", "order_id")
);
CREATE INDEX "order_seller_deleted_at_idx" ON "order"("deleted_at");

CREATE TABLE "order_item" (
  "id" UUID PRIMARY KEY,
  "price" NUMERIC NOT NULL DEFAULT 0,
  "value" NUMERIC NOT NULL,
  "total_price" NUMERIC NOT NULL DEFAULT 0,
  "total_discount_price" NUMERIC NOT NULL DEFAULT 0,
  "custom_discount_type" order_discount_type NOT NULL DEFAULT '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd',
  "custom_discount_value" NUMERIC NOT NULL DEFAULT 0,
  "seller_id" UUID,
  "order_id" UUID NOT NULL REFERENCES "order"("id") ON DELETE CASCADE,
  "product_id" UUID NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_by" UUID,
  "deleted_at" BIGINT NOT NULL DEFAULT 0,
  "deleted_by" UUID,
  UNIQUE ("order_id", "product_id", "deleted_at")
);
CREATE INDEX "order_item_deleted_at_idx" ON "order_item"("deleted_at");
CREATE INDEX "order_item_order_id_idx" ON "order_item"("order_id");

CREATE TABLE "transaction" (
  "id" UUID PRIMARY KEY,
  "value" NUMERIC NOT NULL,
  "order_id" UUID NOT NULL REFERENCES "order"("id") ON DELETE CASCADE,
  "payment_type_id" UUID NOT NULL,
  "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
  "created_by" UUID,
  "deleted_at" BIGINT NOT NULL DEFAULT 0,
  "deleted_by" UUID,
  UNIQUE ("order_id", "payment_type_id", "deleted_at")
);
CREATE INDEX "transaction_deleted_at_idx" ON "transaction"("deleted_at");
CREATE INDEX "transaction_order_id_idx" ON "transaction"("order_id");


-- functions
CREATE OR REPLACE FUNCTION create_order_external_id()
  RETURNS TRIGGER
  LANGUAGE PLPGSQL
  AS
$$
DECLARE
  external_id INT := 0;

BEGIN
    SELECT
      COUNT(*) AS total
    FROM
      "order"
    WHERE
      "shop_id" = NEW."shop_id"
    INTO
      external_id;

    external_id := external_id + 1;

    NEW."external_id"=RIGHT(CONCAT('000000', external_id), 6);

    RETURN NEW;
END; 
$$;

-- triggers
CREATE OR REPLACE TRIGGER create_order_external_id
    BEFORE INSERT ON "order"
    FOR EACH ROW
    EXECUTE PROCEDURE create_order_external_id();

-- functions
CREATE OR REPLACE FUNCTION update_order()
  RETURNS TRIGGER
  LANGUAGE PLPGSQL
  AS
$$

BEGIN

  UPDATE
		"order"
	SET
		total_price = subquery.total_price,
    total_discount_price = subquery.total_discount_price,
    product_sort_count = subquery.product_sort_count
	FROM (
    SELECT
      COALESCE(SUM(total_price), 0) AS total_price,
      COALESCE(SUM(total_discount_price), 0) AS total_discount_price,
      COUNT(*) AS product_sort_count
    FROM "order_item" oi
    WHERE oi.order_id = NEW."order_id" AND oi.deleted_at = 0
  ) AS subquery
	WHERE
		id = NEW."order_id" AND deleted_at = 0;

  RETURN NEW;
END;
$$;

-- triggers
CREATE OR REPLACE TRIGGER update_order
    AFTER INSERT OR UPDATE ON "order_item"
    FOR EACH ROW
    EXECUTE PROCEDURE update_order();

-- functions
CREATE OR REPLACE FUNCTION update_order_item()
  RETURNS TRIGGER
  LANGUAGE PLPGSQL
  AS
$$
DECLARE
  price NUMERIC := 0;

BEGIN

  SELECT sh.retail_price
  FROM "order" o
  JOIN "shop_price" sh ON sh."shop_id" = o."shop_id" AND sh."product_id" = NEW."product_id"
  WHERE o."id" = NEW."order_id" AND o."deleted_at" = 0
  INTO price;

  price = COALESCE(price, 0);

  NEW."price" = price;
  NEW."total_price" = price * NEW."value";

  IF NEW."custom_discount_type" = '9a2aa8fe-806e-44d7-8c9d-575fa67ebefd'::order_discount_type THEN -- none

    NEW."total_discount_price" := 0;

  ELSIF NEW."custom_discount_type" = '9fb3ada6-a73b-4b81-9295-5c1605e54552'::order_discount_type THEN -- amount

    NEW."total_discount_price" := NEW."custom_discount_value";

  ELSIF NEW."custom_discount_type" = '1fe92aa8-2a61-4bf1-b907-182b497584ad'::order_discount_type THEN -- percentage

    NEW."total_discount_price" := NEW."total_price" * NEW."custom_discount_value" / 100;

  END IF;

  RETURN NEW;
END;
$$;

-- triggers
CREATE OR REPLACE TRIGGER update_order_item
    BEFORE INSERT OR UPDATE ON "order_item"
    FOR EACH ROW
    EXECUTE PROCEDURE update_order_item();

-- functions
CREATE OR REPLACE FUNCTION update_order_on_create_transaction()
    RETURNS TRIGGER
    LANGUAGE PLPGSQL
AS
$$
BEGIN

    UPDATE "order"
    SET payed_time = NEW."created_at"
    WHERE id = NEW."order_id" AND deleted_at = 0;

    RETURN NEW;
END;
$$;

-- triggers
CREATE OR REPLACE TRIGGER update_order_on_create_transaction
    AFTER INSERT
    ON "transaction"
    FOR EACH ROW
EXECUTE PROCEDURE update_order_on_create_transaction();
