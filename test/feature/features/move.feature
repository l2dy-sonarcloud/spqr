Feature: Move test
  Background:
    #
    # Make host "coordinator" take control
    #
    Given cluster is up and running
    And host "coordinator2" is stopped
    And host "coordinator2" is started
    
    When I execute SQL on host "coordinator"
    """
    CREATE DISTRIBUTION ds1 COLUMN TYPES integer;
    ADD KEY RANGE krid1 FROM 1 ROUTE TO sh1 FOR DISTRIBUTION ds1;
    ADD KEY RANGE krid2 FROM 11 ROUTE TO sh2 FOR DISTRIBUTION ds1;
    ALTER DISTRIBUTION ds1 ATTACH RELATION xMove DISTRIBUTION KEY w_id;
    ALTER DISTRIBUTION ds1 ATTACH RELATION xMove2 DISTRIBUTION KEY w_id;
    """
    Then command return code should be "0"

  Scenario: MOVE KEY RANGE works
    When I run SQL on host "shard1"
    """
    CREATE TABLE xMove(w_id INT, s TEXT);
    insert into xMove(w_id, s) values(1, '001');
    """
    Then command return code should be "0"
    When I run SQL on host "shard2"
    """
    CREATE TABLE xMove(w_id INT, s TEXT);
    insert into xMove(w_id, s) values(11, '002');
    """
    Then command return code should be "0"
    When I run SQL on host "shard2"
    """
    SELECT * FROM xMove
    """
    Then command return code should be "0"
    And SQL result should match regexp
    """
    002
    """
    When I run SQL on host "shard1"
    """
    SELECT * FROM xMove
    """
    Then command return code should be "0"
    And SQL result should match regexp
    """
    001
    """
    When I execute SQL on host "coordinator"
    """
    MOVE KEY RANGE krid1 to sh2
    """
    Then command return code should be "0"
    When I run SQL on host "shard2"
    """
    SELECT * FROM xMove
    """
    Then command return code should be "0"
    And SQL result should match regexp
    """
    .*002(.|\n)*001
    """
    When I run SQL on host "shard1"
    """
    SELECT * FROM xMove
    """
    Then command return code should be "0"
    And SQL result should not match regexp
    """
    001
    """
    When I run SQL on host "coordinator"
    """
    SHOW key_ranges
    """
    Then command return code should be "0"
    And SQL result should match json_exactly
    """
    [
      {
      "Key range ID":"krid1",
      "Distribution ID":"ds1",
      "Lower bound":"1",
      "Shard ID":"sh2"
      },
      {
      "Key range ID":"krid2",
      "Distribution ID":"ds1",
      "Lower bound":"11",
      "Shard ID":"sh2"
      }
    ]
    """

  Scenario: MOVE KEY RANGE works with varchar keys
    When I execute SQL on host "coordinator"
    """
    CREATE DISTRIBUTION ds2 COLUMN TYPES varchar;
    ADD KEY RANGE krid3 FROM a ROUTE TO sh1 FOR DISTRIBUTION ds2;
    ADD KEY RANGE krid4 FROM aa ROUTE TO sh2 FOR DISTRIBUTION ds2;
    ALTER DISTRIBUTION ds2 ATTACH RELATION xMoveStr DISTRIBUTION KEY w_id;
    """
    Then command return code should be "0"
    When I run SQL on host "shard1"
    """
    CREATE TABLE xMoveStr(w_id TEXT, s TEXT);
    insert into xMoveStr(w_id, s) values('a', '001');
    """
    Then command return code should be "0"
    When I run SQL on host "shard2"
    """
    CREATE TABLE xMoveStr(w_id TEXT, s TEXT);
    insert into xMoveStr(w_id, s) values('aa', '002');
    """
    Then command return code should be "0"
    When I run SQL on host "shard2"
    """
    SELECT * FROM xMoveStr
    """
    Then command return code should be "0"
    And SQL result should match regexp
    """
    002
    """
    When I run SQL on host "shard1"
    """
    SELECT * FROM xMoveStr
    """
    Then command return code should be "0"
    And SQL result should match regexp
    """
    001
    """
    When I execute SQL on host "coordinator"
    """
    MOVE KEY RANGE krid3 to sh2
    """
    Then command return code should be "0"
    When I run SQL on host "shard2"
    """
    SELECT * FROM xMoveStr
    """
    Then command return code should be "0"
    And SQL result should match regexp
    """
    .*002(.|\n)*001
    """
    When I run SQL on host "shard1"
    """
    SELECT * FROM xMoveStr
    """
    Then command return code should be "0"
    And SQL result should not match regexp
    """
    001
    """
    When I run SQL on host "coordinator"
    """
    SHOW key_ranges
    """
    Then command return code should be "0"
    And SQL result should match json_exactly
    """
    [
      {
      "Key range ID":"krid1",
      "Distribution ID":"ds1",
      "Lower bound":"1",
      "Shard ID":"sh1"
      },
      {
      "Key range ID":"krid2",
      "Distribution ID":"ds1",
      "Lower bound":"11",
      "Shard ID":"sh2"
      },
      {
      "Key range ID":"krid3",
      "Distribution ID":"ds2",
      "Lower bound":"'a'",
      "Shard ID":"sh2"
      },
      {
      "Key range ID":"krid4",
      "Distribution ID":"ds2",
      "Lower bound":"'aa'",
      "Shard ID":"sh2"
      }
    ]
    """

  Scenario: MOVE KEY RANGE works with many rows
    When I run SQL on host "shard1"
    """
    CREATE TABLE xMove(w_id INT, s TEXT);
    insert into xMove(w_id, s) values(1, '001');
    insert into xMove(w_id, s) values(2, '002');
    insert into xMove(w_id, s) values(3, '003');
    insert into xMove(w_id, s) values(4, '004');
    insert into xMove(w_id, s) values(5, '005');
    """
    Then command return code should be "0"
    When I run SQL on host "shard2"
    """
    CREATE TABLE xMove(w_id INT, s TEXT);
    """
    Then command return code should be "0"
    When I execute SQL on host "coordinator"
    """
    MOVE KEY RANGE krid1 to sh2
    """
    Then command return code should be "0"
    When I run SQL on host "shard2"
    """
    SELECT * FROM xMove
    """
    Then command return code should be "0"
    And SQL result should match regexp
    """
    .*001(.|\n)*002(.|\n)*003(.|\n)*004(.|\n)*005
    """
    When I run SQL on host "shard1"
    """
    SELECT * FROM xMove
    """
    Then command return code should be "0"
    And SQL result should not match regexp
    """
    .*001(.|\n)*002(.|\n)*003(.|\n)*004(.|\n)*005
    """
    When I run SQL on host "coordinator"
    """
    SHOW key_ranges
    """
    Then command return code should be "0"
    And SQL result should match json_exactly
    """
    [
      {
      "Key range ID":"krid1",
      "Distribution ID":"ds1",
      "Lower bound":"1",
      "Shard ID":"sh2"
      },
      {
      "Key range ID":"krid2",
      "Distribution ID":"ds1",
      "Lower bound":"11",
      "Shard ID":"sh2"
      }
    ]
    """

  Scenario: MOVE KEY RANGE works with many tables
    When I run SQL on host "shard1"
    """
    CREATE TABLE xMove(w_id INT, s TEXT);
    insert into xMove(w_id, s) values(1, '001');
    CREATE TABLE xMove2(w_id INT, s TEXT);
    insert into xMove2(w_id, s) values(2, '002');
    """
    Then command return code should be "0"
    When I run SQL on host "shard2"
    """
    CREATE TABLE xMove(w_id INT, s TEXT);
    CREATE TABLE xMove2(w_id INT, s TEXT);
    """
    Then command return code should be "0"
    When I execute SQL on host "coordinator"
    """
    MOVE KEY RANGE krid1 to sh2
    """
    Then command return code should be "0"
    When I run SQL on host "shard2"
    """
    SELECT * FROM xMove
    """
    Then command return code should be "0"
    And SQL result should match regexp
    """
    001
    """
    When I run SQL on host "shard2"
    """
    SELECT * FROM xMove2
    """
    Then command return code should be "0"
    And SQL result should match regexp
    """
    002
    """
    When I run SQL on host "shard1"
    """
    SELECT * FROM xMove
    """
    Then command return code should be "0"
    And SQL result should not match regexp
    """
    001
    """
    When I run SQL on host "shard1"
    """
    SELECT * FROM xMove2
    """
    Then command return code should be "0"
    And SQL result should not match regexp
    """
    002
    """
    When I run SQL on host "coordinator"
    """
    SHOW key_ranges
    """
    Then command return code should be "0"
    And SQL result should match json_exactly
    """
    [
      {
      "Key range ID":"krid1",
      "Distribution ID":"ds1",
      "Lower bound":"1",
      "Shard ID":"sh2"
      },
      {
      "Key range ID":"krid2",
      "Distribution ID":"ds1",
      "Lower bound":"11",
      "Shard ID":"sh2"
      }
    ]
    """

  Scenario: Move to nonexistent shard fails
    When I run SQL on host "coordinator"
    """
    MOVE KEY RANGE krid1 TO nonexistent
    """
    Then command return code should be "1"
    And SQL error on host "coordinator" should match regexp
    """
    shard with ID .* not found in config
    """

  Scenario: Move nonexistent key range fails
    When I run SQL on host "coordinator"
    """
    MOVE KEY RANGE krid3 TO sh2
    """
    Then command return code should be "1"
    And SQL error on host "coordinator" should match regexp
    """
    no key range found at /keyranges/krid3
    """

  Scenario: Move fails when table does not exist on receiver
    When I run SQL on host "shard1"
    """
    CREATE TABLE xMove(w_id INT, s TEXT);
    insert into xMove(w_id, s) values(1, '001');
    """
    Then command return code should be "0"
    When I run SQL on host "shard2"
    """
    SELECT * FROM xMove
    """
    Then command return code should be "1"
    And SQL error on host "shard2" should match regexp
    """
    relation .* does not exist
    """
    When I run SQL on host "shard1"
    """
    SELECT * FROM xMove
    """
    Then command return code should be "0"
    And SQL result should match regexp
    """
    001
    """
    When I run SQL on host "coordinator"
    """
    MOVE KEY RANGE krid1 to sh2
    """
    Then command return code should be "1"
    And SQL error on host "coordinator" should match regexp
    """
    relation xMove does not exist on receiving shard
    """