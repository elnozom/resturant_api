
GO
CREATE  PROCEDURE SettingsUpdate(
	@id INT,
	@value nvarchar(128)
)
AS
BEGIN
    UPDATE Settings SET SettingValue = @value WHERE SettingID = @id
END

