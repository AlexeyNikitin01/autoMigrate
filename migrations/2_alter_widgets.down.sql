-- скрипты отката к предыдущей версии
ALTER TABLE IF EXISTS widgets
    DROP COLUMN IF EXISTS quantity;
