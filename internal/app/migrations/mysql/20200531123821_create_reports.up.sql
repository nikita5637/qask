CREATE TABLE IF NOT EXISTS `reports` (
`id` bigint(20) NOT NULL,
  `user_id` bigint(20) NOT NULL,
  `message` varchar(1024) NOT NULL,
  `date` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
) ENGINE=InnoDB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

ALTER TABLE `reports`
 ADD PRIMARY KEY (`id`), ADD KEY `user_id` (`user_id`);

ALTER TABLE `reports`
MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT,AUTO_INCREMENT=1;

ALTER TABLE `reports`
ADD CONSTRAINT `reports_ibfk_1` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);