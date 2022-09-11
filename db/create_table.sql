DROP TABLE IF EXISTS user;
CREATE TABLE user(
    id bigint NOT NULL AUTO_INCREMENT  COMMENT '主键' ,
    name VARCHAR(255) NOT NULL   COMMENT '用户名' ,
    login_name VARCHAR(255) NOT NULL   COMMENT '登录名' ,
    `password` VARCHAR(255) NOT NULL   COMMENT '登录密码,sha256加密' ,
    create_time DATETIME NOT NULL  DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间' ,
    email VARCHAR(255) NOT NULL   COMMENT '邮箱' ,
    phone_number VARCHAR(255) NOT NULL   COMMENT '手机号码' ,
    PRIMARY KEY (id)
)  COMMENT = '用户表';

CREATE UNIQUE INDEX user_idx_login_name ON user(login_name);
CREATE UNIQUE INDEX user_idx_email ON user(email);
CREATE UNIQUE INDEX user_idx_phone_number ON user(phone_number);


DROP TABLE IF EXISTS node;
CREATE TABLE node(
    id bigint NOT NULL AUTO_INCREMENT  COMMENT '主键' ,
    host VARCHAR(255) NOT NULL   COMMENT '主机地址 IP或者域名' ,
    port VARCHAR(255) NOT NULL   COMMENT 'exporter对应的端口号' ,
    owner VARCHAR(255)    COMMENT '所有者/责任人' ,
    job_name VARCHAR(255)    COMMENT 'prometheus job_name' ,
    `group` VARCHAR(255) NOT NULL   COMMENT '组名' ,
    labels VARCHAR(255)    COMMENT '标签 对应 prometheus中的label配置选项' ,
    create_time DATETIME   DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间' ,
    PRIMARY KEY (id)
)  COMMENT = '主机';


CREATE UNIQUE INDEX node_idx_hosts ON node(host,port);

DROP TABLE IF EXISTS `group`;
CREATE TABLE `group`(
    id bigint NOT NULL AUTO_INCREMENT  COMMENT '主键' ,
    name VARCHAR(255) NOT NULL   COMMENT '分组名称' ,
    identification VARCHAR(255) NOT NULL   COMMENT '英文标识' ,
    create_time DATETIME   DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间' ,
    PRIMARY KEY (id)
)  COMMENT = '分组';


CREATE UNIQUE INDEX group_idx_name ON `group`(name);

DROP TABLE IF EXISTS namespace;
CREATE TABLE namespace(
    id bigint NOT NULL AUTO_INCREMENT  COMMENT '主键' ,
    name VARCHAR(255) NOT NULL   COMMENT '名称' ,
    identification VARCHAR(255) NOT NULL   COMMENT '标识' ,
    create_time DATETIME    COMMENT '创建时间' ,
    PRIMARY KEY (id)
)  COMMENT = '命名空间';


CREATE UNIQUE INDEX namespace_idx_name ON namespace(name);
CREATE INDEX namespace_idx_identification ON namespace(identification);

