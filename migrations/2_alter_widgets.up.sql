-- скрипты создания сущностей
ALTER TABLE IF EXISTS widgets
 ADD COLUMN quantity integer DEFAULT 0 NOT NULL;
