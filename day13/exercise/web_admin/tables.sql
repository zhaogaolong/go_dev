
# app tables
create table tbl_app_info(
    app_id int auto_increment primary key,
    app_name varchar(1024) not null,
    app_type varchar(64) not null,
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    develop_path varchar(1024) not null 
) ENGINE=InnoDB DEFAULT CHARSET=utf8 auto_increment=1;


create table tbl_app_ip (
    app_id int,
    ip varchar(64),
    Key app_id_ip_index (app_id, ip)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 auto_increment=1;

# log tables
create table tbl_log_info(
    log_id int auto_increment primary key,
    app_id varchar(1024) not null,
    create_time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    log_path varchar(1024) not null,
    status TINYINT default 1,
    topic varchar(1024)
) ENGINE=InnoDB DEFAULT CHARSET=utf8 auto_increment=1;

