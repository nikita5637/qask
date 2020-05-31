CREATE TABLE IF NOT EXISTS `users` (
    `id` bigint(20) NOT NULL,
    `tgid` bigint(20) DEFAULT NULL,
    `username` varchar(100) NOT NULL,
    `firstname` varchar(100) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

ALTER TABLE `users`
    ADD PRIMARY KEY (`id`), ADD UNIQUE KEY `username` (`username`);

ALTER TABLE `users`
    MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT;