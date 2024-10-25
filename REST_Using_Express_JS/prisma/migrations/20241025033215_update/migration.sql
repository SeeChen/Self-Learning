-- RedefineTables
PRAGMA defer_foreign_keys=ON;
PRAGMA foreign_keys=OFF;
CREATE TABLE "new_UserCompany" (
    "userId" INTEGER NOT NULL,
    "companyId" INTEGER NOT NULL,
    "role" INTEGER NOT NULL,

    PRIMARY KEY ("userId", "companyId"),
    CONSTRAINT "UserCompany_userId_fkey" FOREIGN KEY ("userId") REFERENCES "User" ("id") ON DELETE CASCADE ON UPDATE CASCADE,
    CONSTRAINT "UserCompany_companyId_fkey" FOREIGN KEY ("companyId") REFERENCES "Company" ("id") ON DELETE CASCADE ON UPDATE CASCADE
);
INSERT INTO "new_UserCompany" ("companyId", "role", "userId") SELECT "companyId", "role", "userId" FROM "UserCompany";
DROP TABLE "UserCompany";
ALTER TABLE "new_UserCompany" RENAME TO "UserCompany";
CREATE UNIQUE INDEX "UserCompany_userId_key" ON "UserCompany"("userId");
PRAGMA foreign_keys=ON;
PRAGMA defer_foreign_keys=OFF;
