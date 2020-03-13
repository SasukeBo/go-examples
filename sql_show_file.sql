SELECT COUNT(p.id) AS total, p.store_id
FROM module_builder_3_new_permission AS p
LEFT JOIN module_builder_3_store_installs AS store_installs ON store_installs.deleted_at IS NULL AND store_installs.app_id = p.app_id
LEFT JOIN module_builder_3_apps AS apps ON apps.deleted_at IS NULL AND apps.id = store_installs.app_id
LEFT JOIN stores ON stores.deleted != 1 AND stores.id = p.store_id
WHERE p.store_id IN ('36091','2221') AND p.deleted_at IS NULL GROUP BY p.store_id;


SELECT p.*,apps.name AS name,apps.name AS appid, stores.title AS __store_id$display, apps.name AS __app_id$display
FROM module_builder_3_new_permission AS p
LEFT JOIN module_builder_3_store_installs AS store_installs ON store_installs.deleted_at IS NULL AND store_installs.app_id = p.app_id
LEFT JOIN module_builder_3_apps AS apps ON apps.deleted_at IS NULL AND apps.id = store_installs.app_id
LEFT JOIN stores ON stores.deleted != 1 AND stores.id = p.store_id
WHERE p.store_id IN (36091,2221) AND p.deleted_at IS NULL
ORDER BY p.id DESC LIMIT 50 OFFSET 0;

CREATE TABLE `module_builder_3_audit_logs` (`id` BIGINT(20) unsigned,`store_id` int(11) ,`user_id` int(11),`type` varchar(255),`detail` varchar(255),`request_id` varchar(255),`created_at` datetime , PRIMARY KEY (`id`));

DELIMITER $
CREATE PROCEDURE  pro_clear_data()
BEGIN
    SET @news_date=(select UNIX_TIMESTAMP((select date_sub(NOW(),interval 1 MONTH))) * 1000);
    /*设定删除数据的时间*/
    SET @max_id=(SELECT MAX(messageID) from sys_msg_offline where creationDate <= @news_date);
    /*设定需要删除数据的最大messageID*/
    SET @news_count=(SELECT COUNT(messageID) FROM sys_msg_offline where messageID <= @max_id);
    /*设定删除数据数量*/
    WHILE @news_count >0 DO
    /*删除需要删除的数据，一次最多1000*/
    SET @news_count=(SELECT COUNT(messageID) FROM sys_msg_offline where messageID <= @max_id);
    /*刷新删除数据数量*/
    END WHILE;
END $
