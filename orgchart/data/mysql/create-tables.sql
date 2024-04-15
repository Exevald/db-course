CREATE TABLE branch
(
    `branch_id` BINARY(16) NOT NULL,
    `city`      VARCHAR(100) NOT NULL,
    `address`   VARCHAR(255) NOT NULL,
    PRIMARY KEY (`branch_id`),
    UNIQUE (`city`, `address`)
)
    ENGINE = InnoDB
    CHARACTER SET = utf8mb4
    COLLATE utf8mb4_unicode_ci
;

CREATE TABLE employee
(
    `employee_id` BINARY(16) NOT NULL,
    `branch_id`   BINARY(16) NOT NULL,
    `first_name`  VARCHAR(255) NOT NULL,
    `last_name`   VARCHAR(255) NOT NULL,
    `middle_name` VARCHAR(255) NOT NULL,
    `job_title`   VARCHAR(255) NOT NULL,
    `phone`       VARCHAR(255) NOT NULL,
    `email`       VARCHAR(255) NOT NULL,
    `gender`      TINYINT(1)   NOT NULL,
    `birth_date`  DATE         NOT NULL,
    `hire_date`   DATE         NOT NULL,
    `comment`     TEXT,
    `avatar_path` TEXT,
    PRIMARY KEY (`employee_id`),
    FOREIGN KEY (`branch_id`) REFERENCES branch (`branch_id`) ON DELETE CASCADE
)
    ENGINE = InnoDB
    CHARACTER SET = utf8mb4
    COLLATE utf8mb4_unicode_ci
;