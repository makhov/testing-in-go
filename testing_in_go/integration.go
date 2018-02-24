// FLAG OMIT
// +build integration
// ENDFLAG OMIT

// DOC OMIT
// Этот пакет содержит функциональные интеграционные тесты, 
// тестирующие приложение по методу черного ящика.
// Для его работы необходимы:
// …
package integration_test
// END OMIT

// FLAGRUN OMIT
go test -tags=integration
// END OMIT