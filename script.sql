create table sliver_comments
(
    ID              bigint auto_increment
        primary key,
    wallet_address  varchar(128) null,
    comment_content text         null,
    parent_id       bigint       null,
    p_id            bigint       null,
    created_at      timestamp    null
);

create table sliver_plan
(
    ID              bigint auto_increment
        primary key,
    p_id            bigint          null,
    client_address  varchar(256)    null,
    data_cap        decimal(36, 18) null,
    staking_amount  decimal(36, 18) null,
    staking_id      bigint          null,
    staking_days    int             null,
    staking_address varchar(256)    null,
    status          varchar(52)     null,
    allocate_time   timestamp       null,
    allocate_tx     varchar(256)    null,
    created_at      timestamp       null,
    staking_time    timestamp       null
);

create table sliver_proposals
(
    ID                    bigint auto_increment comment '唯一自增ID'
        primary key,
    p_id                  varchar(128)                     null comment '项目ID',
    client_address        varchar(256)                     null,
    p_name                varchar(128)                     null comment '项目名称',
    p_content             json                             null comment '项目内容',
    p_user                bigint                           null comment '用户ID',
    status                varchar(32)                      null comment '状态值',
    reason_rejection      text                             null,
    request_data_cap      varchar(52)                      null,
    data_cap              varchar(128)                     null comment '申领份额',
    kyc_status            varchar(32) default 'unverified' null,
    kyc_verification_time timestamp                        null,
    created_at            timestamp                        null comment '创建时间',
    update_at             timestamp                        null comment '修改时间',
    constraint sliver_proposals_pk
        unique (p_id)
);

create table sliver_user
(
    ID            bigint auto_increment comment '唯一自增ID'
        primary key,
    login_name    varchar(128) null comment 'login name',
    login_pass    varchar(512) null,
    role          varchar(32)  null comment 'role',
    wallet        varchar(128) null comment '钱包地址',
    display_name  varchar(128) null comment '显示名称',
    email         varchar(128) null comment '邮箱地址',
    kyc_status    varchar(32)  null comment 'kyc状态',
    created_at    timestamp    null comment '创建时间',
    last_login_at timestamp    null comment '最后登录时间',
    constraint sliver_user_pk_wallet
        unique (wallet)
);


