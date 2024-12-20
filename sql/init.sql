CREATE TABLE public."Users" (
    "Id" UUID NOT NULL DEFAULT gen_random_uuid(),
    "UserName" VARCHAR(50) NOT NULL,
    "Password" TEXT NOT NULL,
    "CreatedDate" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "IsAdmin" BOOLEAN NOT NULL DEFAULT FALSE
);