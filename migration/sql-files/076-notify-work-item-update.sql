CREATE OR REPLACE FUNCTION notify_work_item_update() RETURNS TRIGGER AS $$
BEGIN
        -- Execute pg_notify(channel, notification)
        PERFORM pg_notify('wi_update', NEW.id::text);
        
        -- Result is ignored since this is an AFTER trigger
        RETURN NULL; 
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER work_item_update_notification_trigger
AFTER UPDATE ON work_items
    FOR EACH ROW EXECUTE PROCEDURE notify_work_item_update();
