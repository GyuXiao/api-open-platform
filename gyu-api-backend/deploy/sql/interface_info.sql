use api_db;

-- 接口信息
create table if not exists api_db.`interfaceInfo`
(
`id` bigint not null auto_increment comment '主键' primary key,
`name` varchar(128) not null comment '接口名称',
`description` varchar(256) null comment '接口描述',
`url` varchar(512) not null comment '接口地址',
`requestParams` text null comment '请求参数',
`requestHeader` text null comment '请求头',
`responseHeader` text null comment '响应头',
`status` int not null comment '接口状态（0-关闭 1-开启）',
`method` varchar(256) not null comment '请求类型',
`userId` bigint not null comment '用户 Id',
`createTime` datetime default CURRENT_TIMESTAMP not null comment '创建时间',
`updateTime` datetime default CURRENT_TIMESTAMP not null on update CURRENT_TIMESTAMP comment '更新时间',
`isDelete` tinyint default 0 not null comment '是否删除(0-未删, 1-已删)'
) comment '接口信息';

-- 插入初始数据
insert into api_db.`interfaceInfo` (`name`, `description`, `url`, `requestHeader`, `responseHeader`, `status`, `method`, `userId`) values ('谭文轩', '王航', 'www.chance-kutch.org', '冯昊然', '孔鹤轩', 0, '覃熠彤', 4497220944);
insert into api_db.`interfaceInfo` (`name`, `description`, `url`, `requestHeader`, `responseHeader`, `status`, `method`, `userId`) values ('姚鑫鹏', '赵天磊', 'www.eldon-abshire.io', '阎明', '任浩然', 0, '贾天磊', 653);
insert into api_db.`interfaceInfo` (`name`, `description`, `url`, `requestHeader`, `responseHeader`, `status`, `method`, `userId`) values ('王笑愚', '周健柏', 'www.frances-lockman.io', '贺思远', '贺果', 0, '郑天宇', 5773583);
insert into api_db.`interfaceInfo` (`name`, `description`, `url`, `requestHeader`, `responseHeader`, `status`, `method`, `userId`) values ('薛炫明', '周皓轩', 'www.sheridan-erdman.com', '魏天磊', '孙峻熙', 0, '丁子涵', 756607);
insert into api_db.`interfaceInfo` (`name`, `description`, `url`, `requestHeader`, `responseHeader`, `status`, `method`, `userId`) values ('邹明辉', '戴烨霖', 'www.shannon-predovic.biz', '宋子骞', '黎明杰', 0, '叶乐驹', 4);
insert into api_db.`interfaceInfo` (`name`, `description`, `url`, `requestHeader`, `responseHeader`, `status`, `method`, `userId`) values ('罗睿渊', '周耀杰', 'www.isa-conroy.biz', '杨懿轩', '覃旭尧', 0, '叶思源', 908);
insert into api_db.`interfaceInfo` (`name`, `description`, `url`, `requestHeader`, `responseHeader`, `status`, `method`, `userId`) values ('徐嘉懿', '蔡伟宸', 'www.debbi-mosciski.io', '苏锦程', '黎越彬', 0, '汪思', 6018);
insert into api_db.`interfaceInfo` (`name`, `description`, `url`, `requestHeader`, `responseHeader`, `status`, `method`, `userId`) values ('任鑫磊', '孙绍齐', 'www.britni-romaguera.co', '叶明辉', '钟绍齐', 0, '白思淼', 4841);
insert into api_db.`interfaceInfo` (`name`, `description`, `url`, `requestHeader`, `responseHeader`, `status`, `method`, `userId`) values ('钟志强', '薛驰', 'www.maybelle-pfeffer.biz', '陈正豪', '陈智宸', 0, '姚锦程', 59472549);
insert into api_db.`interfaceInfo` (`name`, `description`, `url`, `requestHeader`, `responseHeader`, `status`, `method`, `userId`) values ('罗弘文', '田博超', 'www.carmel-turner.info', '于鑫鹏', '彭风华', 0, '谢文博', 5);
insert into api_db.`interfaceInfo` (`name`, `description`, `url`, `requestHeader`, `responseHeader`, `status`, `method`, `userId`) values ('方文轩', '江雨泽', 'www.quinton-corwin.name', '陆明', '金昊焱', 0, '董远航', 6133);
insert into api_db.`interfaceInfo` (`name`, `description`, `url`, `requestHeader`, `responseHeader`, `status`, `method`, `userId`) values ('贺伟祺', '龙胤祥', 'www.sixta-hansen.name', '蔡峻熙', '罗耀杰', 0, '许君浩', 6301070);
insert into api_db.`interfaceInfo` (`name`, `description`, `url`, `requestHeader`, `responseHeader`, `status`, `method`, `userId`) values ('钱晟睿', '雷哲瀚', 'www.keren-mayer.org', '高立果', '董博超', 0, '彭琪', 8287453);
insert into api_db.`interfaceInfo` (`name`, `description`, `url`, `requestHeader`, `responseHeader`, `status`, `method`, `userId`) values ('贺立诚', '何越彬', 'www.wilbert-bogisich.net', '周烨华', '严立轩', 0, '杜健雄', 23834);
insert into api_db.`interfaceInfo` (`name`, `description`, `url`, `requestHeader`, `responseHeader`, `status`, `method`, `userId`) values ('龚致远', '何果', 'www.cleta-mohr.co', '张擎苍', '宋浩然', 0, '熊苑博', 595645);
insert into api_db.`interfaceInfo` (`name`, `description`, `url`, `requestHeader`, `responseHeader`, `status`, `method`, `userId`) values ('陶智渊', '徐伟诚', 'www.eleanora-monahan.name', '王楷瑞', '徐智宸', 0, '赵天宇', 8521793934);
insert into api_db.`interfaceInfo` (`name`, `description`, `url`, `requestHeader`, `responseHeader`, `status`, `method`, `userId`) values ('武君浩', '钱雨泽', 'www.tommie-erdman.com', '雷雨泽', '白昊强', 0, '何修杰', 85);
insert into api_db.`interfaceInfo` (`name`, `description`, `url`, `requestHeader`, `responseHeader`, `status`, `method`, `userId`) values ('傅志强', '钱梓晨', 'www.alexander-stehr.com', '白天宇', '方晓啸', 0, '黎思聪', 4321);
insert into api_db.`interfaceInfo` (`name`, `description`, `url`, `requestHeader`, `responseHeader`, `status`, `method`, `userId`) values ('贾烨霖', '马志泽', 'www.cher-marvin.name', '丁鑫磊', '萧鹏涛', 0, '冯远航', 568990039);
insert into api_db.`interfaceInfo` (`name`, `description`, `url`, `requestHeader`, `responseHeader`, `status`, `method`, `userId`) values ('郝煜城', '黄瑞霖', 'www.jerrell-kiehn.io', '汪昊天', '赵文博', 0, '陆弘文', 7668618995);

ALTER TABLE `interfaceInfo` ADD INDEX `name` (`name`);
