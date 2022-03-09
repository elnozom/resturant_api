

GO
EXEC DropProcIfExist @Name = "CartCallCreate"
GO
CREATE PROC CartCallCreate
    (@CallType BIT ,
    @CartSerial INT ,
    @TableSerial INT ,
    @GuestName VARCHAR(100)
    )
AS
BEGIN
    IF EXISTS (SELECT NozCartCallsSerial FROM NozCartCalls WHERE CartSerial = @CartSerial AND CallType = @CallType  AND GuestName = @GuestName AND RespondedAt IS NULL )
    BEGIN
        SELECT 0 AS Created
        RETURN
    END
    IF @CallType = 1 AND NOT EXISTS (
        SELECT * FROM StkTr03 o 
            WHERE  TableSerial = @TableSerial AND 
                    ISNULL(TotalCash,0) = 0 
        )
        BEGIN
            SELECT 0 AS Created
            RETURN
    END
    IF @CallType = 1 AND  EXISTS (
        SELECT * FROM StkTr03 o 
            JOIN  StkTr04 d 
            ON o.Serial = d.HeadSerial
            WHERE  TableSerial = @TableSerial AND 
                    ISNULL(TotalCash,0) = 0  AND d.Printed = 0 
        )
        BEGIN
            SELECT 0 AS Created
            RETURN
        END

    
    INSERT INTO NozCartCalls
        (CallType , CartSerial , TableSerial , GuestName)
    VALUES
        (@CallType , @CartSerial , @TableSerial , @GuestName)
    SELECT 1 AS Created
END


GO
ALTER PROCEDURE [dbo].[GroupCodeListByGroupTypeId](@GroupTypeID int)
AS
BEGIN

	SET NOCOUNT ON;
	SELECT  g.GroupCode , g.GroupName , g.ImagePath FROM  "GroupCode" g WHERE g.GroupTypeID = @GroupTypeID AND g.ShowOnMenu = 1
END


GO
EXEC DropProcIfExist @Name = "CartListCalls"
GO
CREATE PROC CartListCalls
    (@Imei VARCHAR(100))
AS
BEGIN
    SELECT COUNT(c.CartSerial) , c.TableSerial, c.CallType ,t.GroupTableNo , t.TableNo , ISNULL(c.GuestName , '') , gt.GroupTableName   
    FROM NozCartCalls c 
        JOIN
            Tables t ON c.TableSerial = t.Serial 
		JOIN
            GroupTables gt ON gt.GroupTableNo = t.GroupTableNo 
        JOIN
            GTE_Map g ON t.GroupTableNo = g.GTID 
        JOIN
            ComUse com ON g.EmpID = com.UserId
    WHERE com.Imei = @Imei AND RespondedAt IS NULL
	group by c.TableSerial, c.CallType ,t.GroupTableNo , t.TableNo , c.GuestName , gt.GroupTableName   
END



GO
EXEC DropProcIfExist @Name = "CartCheckCalls"
GO
CREATE PROC CartCheckCalls
    (@Imei VARCHAR(100))
AS
BEGIN
    SELECT  COUNT(*) countItems  
    FROM NozCartCalls c 
        JOIN
            Tables t ON c.TableSerial = t.Serial 
        JOIN
            GTE_Map g ON t.GroupTableNo = g.GTID 
        JOIN
            ComUse com ON g.EmpID = com.UserId
    WHERE com.Imei = @Imei AND RespondedAt IS NULL
END




GO
EXEC DropProcIfExist @Name = "CartCallRespond"
GO
CREATE PROC CartCallRespond
    (@Serials VARCHAR(100) ,
    @WaiterCode INT)
AS
BEGIN
    declare @TableSerial INT
    declare I_Serial cursor
	for
	SELECT Split.a.value('.', 'NVARCHAR(MAX)') DATA
    FROM
	  (
		 SELECT CAST('<X>'+REPLACE(@Serials, ',', '</X><X>')+'</X>' AS XML) AS String
       ) AS A
		CROSS APPLY String.nodes('/X') AS Split(a)
		
		open I_Serial
		Fetch Next From I_Serial into @TableSerial 
	while @@FETCH_STATUS = 0
	begin 
        UPDATE NozCartCalls SET RespondedAt = GETDATE() ,  WaiterCode = @WaiterCode WHERE TableSerial = @TableSerial
		Fetch Next From I_Serial into @TableSerial 
    END
    Close I_Serial
		DEALLOCATE   I_Serial
END



ALTER TABLE GroupCode
ADD ImagePath VARCHAR(100); 
UPDATE GroupCode  SET ImagePath = CONCAT(GroupTypeID , '/' , GroupCode , '/' , 'Default.jpg')