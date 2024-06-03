use api_db;

-- 用户调用接口关系表
create table if not exists api_db.`user_interface_info`
(
    `id` bigint not null auto_increment comment '主键' primary key,
  	`userId` bigint not null comment '调用用户 id',
    `interfaceInfoId` bigint not null comment '接口 id',
   	`totalNum` int default 0 not null comment '总调用次数',
  	`leftNum` int default 0 not null comment '剩余调用次数',
  	`status` int default 0 not null comment '0-正常，1-禁用',
  	`createTime` datetime default CURRENT_TIMESTAMP not null comment '创建时间',
    `updateTime` datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
    `isDelete` tinyint default 0 not null comment '是否删除(0-未删, 1-已删)'
) comment '用户调用接口关系';

-- 插入数据
insert into api_db.`user_interface_info` (`id`, `userId`, `interfaceInfoId`, `totalNum`, `leftNum`, `status`, `createTime`, `updateTime`, `isDelete`) values (2, 651, 7941, 6162, 257, 0, '2022-01-16 21:03:11', '2022-05-29 03:29:58', 0);
insert into api_db.`user_interface_info` (`id`, `userId`, `interfaceInfoId`, `totalNum`, `leftNum`, `status`, `createTime`, `updateTime`, `isDelete`) values (3, 1, 6588, 315, 2002, 0, '2022-08-27 22:07:20', '2022-10-11 14:16:11', 0);
insert into api_db.`user_interface_info` (`id`, `userId`, `interfaceInfoId`, `totalNum`, `leftNum`, `status`, `createTime`, `updateTime`, `isDelete`) values (4, 70, 2631981386, 488018326, 43, 0, '2022-03-17 00:49:46', '2022-03-26 01:47:21', 0);
insert into api_db.`user_interface_info` (`id`, `userId`, `interfaceInfoId`, `totalNum`, `leftNum`, `status`, `createTime`, `updateTime`, `isDelete`) values (5, 28, 8108327, 960262, 4, 0, '2022-12-21 03:19:02', '2022-09-28 20:31:20', 0);
insert into api_db.`user_interface_info` (`id`, `userId`, `interfaceInfoId`, `totalNum`, `leftNum`, `status`, `createTime`, `updateTime`, `isDelete`) values (6, 180, 1, 22737, 672446089, 0, '2022-05-22 12:42:52', '2022-04-10 07:51:19', 0);
insert into api_db.`user_interface_info` (`id`, `userId`, `interfaceInfoId`, `totalNum`, `leftNum`, `status`, `createTime`, `updateTime`, `isDelete`) values (7, 48, 33498648, 786958647, 911819, 0, '2022-02-28 16:09:02', '2022-09-28 01:43:46', 0);
insert into api_db.`user_interface_info` (`id`, `userId`, `interfaceInfoId`, `totalNum`, `leftNum`, `status`, `createTime`, `updateTime`, `isDelete`) values (8, 6, 5004826644, 1140, 4697, 0, '2022-01-09 19:42:05', '2022-04-10 13:20:32', 0);
insert into api_db.`user_interface_info` (`id`, `userId`, `interfaceInfoId`, `totalNum`, `leftNum`, `status`, `createTime`, `updateTime`, `isDelete`) values (9, 68116, 2607701877, 196551, 5246, 0, '2022-11-19 10:16:21', '2022-07-12 08:15:04', 0);
insert into api_db.`user_interface_info` (`id`, `userId`, `interfaceInfoId`, `totalNum`, `leftNum`, `status`, `createTime`, `updateTime`, `isDelete`) values (10, 30, 9, 1000, 40, 0, '2022-11-26 20:05:06', '2022-09-24 23:12:49', 0);
insert into api_db.`user_interface_info` (`id`, `userId`, `interfaceInfoId`, `totalNum`, `leftNum`, `status`, `createTime`, `updateTime`, `isDelete`) values (11, 56088, 278, 832, 54, 0, '2022-05-16 02:13:16', '2022-08-07 13:42:05', 0);

