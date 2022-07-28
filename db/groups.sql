USE RMSS


DROP PROC IF EXISTS GroupCodeListHierarchy
GO
CREATE  PROCEDURE GroupCodeListHierarchy
AS
BEGIN
	WITH groups (childrenLength ,  code )
	AS (
		SELECT COUNT(c.parent_code) , p.code , (LEN(p.code) / 2 - 1) groupLevel  FROM GroupCode p JOIN GroupCode c ON p.code = c.parent_code GROUP BY p.code
	)

    SELECT gc.GroupCode ,  gc.GroupName , ISNULL(gc.GroupNameEn  ,gc.GroupName) GroupNameEn, ISNULL(gc.parent_code , '') parent_code , gc.code , g.childrenLength , g.groupLevel 
        FROM GroupCode gc
        LEFT JOIN groups g ON  g.code = gc.code
    ORDER BY g.groupLevel
END



DROP PROC IF EXISTS GroupCodeInsertUpdate
GO
CREATE  PROCEDURE GroupCodeInsertUpdate (
    @groupCode  INT = 0,
    @groupName  VARCHAR(250),
    @groupNameEn  VARCHAR(250),
    @parentCode  VARCHAR(250),
    @code VARCHAR(250)
)
AS
BEGIN
    IF @groupCode = 0
        INSERT GroupCode (GroupName , GroupNameEn , parent_code , code) 
        VALUES (@groupName , @groupNameEn , @parentCode ,@code)
    ELSE
        BEGIN
            UPDATE GroupCode SET GroupName = @groupName ,GroupNameEn = @groupNameEn ,parentCode = @parentCode ,code = @code  WHERE GroupCode = @groupCode
        END
    
    SELECT GroupCode GroupName , GroupNameEn , ISNULL(parent_code , '') , code 
    FROM GroupCode 
    WHERE GroupCode = dbo.ISZERO(@groupCode , SCOPE_IDENTITY())
END
