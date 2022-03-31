CREATE TABLE "posts" (
                         "id" bigserial PRIMARY KEY,
                         "title" varchar NOT NULL,
                         "content" varchar NOT NULL,
                         "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "tags" (
                        "id" bigserial PRIMARY KEY,
                        "name" varchar NOT NULL,
                        "description" varchar NOT NULL,
                        "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "post_tag" (
                            "post_id" bigserial PRIMARY KEY,
                            "tag_id" bigserial
);

ALTER TABLE "post_tag" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");

ALTER TABLE "post_tag" ADD FOREIGN KEY ("tag_id") REFERENCES "tags" ("id");

CREATE INDEX ON "posts" ("id");

CREATE INDEX ON "tags" ("id");