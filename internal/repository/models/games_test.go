// Code generated by SQLBoiler 4.14.2 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/randomize"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testGames(t *testing.T) {
	t.Parallel()

	query := Games()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testGamesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Game{}
	if err = randomize.Struct(seed, o, gameDBTypes, true, gameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Game struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Games().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGamesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Game{}
	if err = randomize.Struct(seed, o, gameDBTypes, true, gameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Game struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := Games().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Games().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGamesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Game{}
	if err = randomize.Struct(seed, o, gameDBTypes, true, gameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Game struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := GameSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := Games().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGamesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Game{}
	if err = randomize.Struct(seed, o, gameDBTypes, true, gameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Game struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := GameExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if Game exists: %s", err)
	}
	if !e {
		t.Errorf("Expected GameExists to return true, but got false.")
	}
}

func testGamesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Game{}
	if err = randomize.Struct(seed, o, gameDBTypes, true, gameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Game struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	gameFound, err := FindGame(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if gameFound == nil {
		t.Error("want a record, got nil")
	}
}

func testGamesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Game{}
	if err = randomize.Struct(seed, o, gameDBTypes, true, gameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Game struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = Games().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testGamesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Game{}
	if err = randomize.Struct(seed, o, gameDBTypes, true, gameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Game struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := Games().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testGamesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	gameOne := &Game{}
	gameTwo := &Game{}
	if err = randomize.Struct(seed, gameOne, gameDBTypes, false, gameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Game struct: %s", err)
	}
	if err = randomize.Struct(seed, gameTwo, gameDBTypes, false, gameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Game struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = gameOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = gameTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Games().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testGamesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	gameOne := &Game{}
	gameTwo := &Game{}
	if err = randomize.Struct(seed, gameOne, gameDBTypes, false, gameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Game struct: %s", err)
	}
	if err = randomize.Struct(seed, gameTwo, gameDBTypes, false, gameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Game struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = gameOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = gameTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Games().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func gameBeforeInsertHook(ctx context.Context, e boil.ContextExecutor, o *Game) error {
	*o = Game{}
	return nil
}

func gameAfterInsertHook(ctx context.Context, e boil.ContextExecutor, o *Game) error {
	*o = Game{}
	return nil
}

func gameAfterSelectHook(ctx context.Context, e boil.ContextExecutor, o *Game) error {
	*o = Game{}
	return nil
}

func gameBeforeUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Game) error {
	*o = Game{}
	return nil
}

func gameAfterUpdateHook(ctx context.Context, e boil.ContextExecutor, o *Game) error {
	*o = Game{}
	return nil
}

func gameBeforeDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Game) error {
	*o = Game{}
	return nil
}

func gameAfterDeleteHook(ctx context.Context, e boil.ContextExecutor, o *Game) error {
	*o = Game{}
	return nil
}

func gameBeforeUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Game) error {
	*o = Game{}
	return nil
}

func gameAfterUpsertHook(ctx context.Context, e boil.ContextExecutor, o *Game) error {
	*o = Game{}
	return nil
}

func testGamesHooks(t *testing.T) {
	t.Parallel()

	var err error

	ctx := context.Background()
	empty := &Game{}
	o := &Game{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, gameDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Game object: %s", err)
	}

	AddGameHook(boil.BeforeInsertHook, gameBeforeInsertHook)
	if err = o.doBeforeInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	gameBeforeInsertHooks = []GameHook{}

	AddGameHook(boil.AfterInsertHook, gameAfterInsertHook)
	if err = o.doAfterInsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	gameAfterInsertHooks = []GameHook{}

	AddGameHook(boil.AfterSelectHook, gameAfterSelectHook)
	if err = o.doAfterSelectHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	gameAfterSelectHooks = []GameHook{}

	AddGameHook(boil.BeforeUpdateHook, gameBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	gameBeforeUpdateHooks = []GameHook{}

	AddGameHook(boil.AfterUpdateHook, gameAfterUpdateHook)
	if err = o.doAfterUpdateHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	gameAfterUpdateHooks = []GameHook{}

	AddGameHook(boil.BeforeDeleteHook, gameBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	gameBeforeDeleteHooks = []GameHook{}

	AddGameHook(boil.AfterDeleteHook, gameAfterDeleteHook)
	if err = o.doAfterDeleteHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	gameAfterDeleteHooks = []GameHook{}

	AddGameHook(boil.BeforeUpsertHook, gameBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	gameBeforeUpsertHooks = []GameHook{}

	AddGameHook(boil.AfterUpsertHook, gameAfterUpsertHook)
	if err = o.doAfterUpsertHooks(ctx, nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	gameAfterUpsertHooks = []GameHook{}
}

func testGamesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Game{}
	if err = randomize.Struct(seed, o, gameDBTypes, true, gameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Game struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Games().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testGamesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Game{}
	if err = randomize.Struct(seed, o, gameDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Game struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(gameColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := Games().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testGameToOneUserUsingOwnerUser(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Game
	var foreign User

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, gameDBTypes, true, gameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Game struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, userDBTypes, false, userColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize User struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	queries.Assign(&local.OwnerUserID, foreign.ID)
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.OwnerUser().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if !queries.Equal(check.ID, foreign.ID) {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	ranAfterSelectHook := false
	AddUserHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *User) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := GameSlice{&local}
	if err = local.L.LoadOwnerUser(ctx, tx, false, (*[]*Game)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.OwnerUser == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.OwnerUser = nil
	if err = local.L.LoadOwnerUser(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.OwnerUser == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testGameToOnePublisherUsingPublisher(t *testing.T) {
	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var local Game
	var foreign Publisher

	seed := randomize.NewSeed()
	if err := randomize.Struct(seed, &local, gameDBTypes, false, gameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Game struct: %s", err)
	}
	if err := randomize.Struct(seed, &foreign, publisherDBTypes, false, publisherColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Publisher struct: %s", err)
	}

	if err := foreign.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	local.PublisherID = foreign.ID
	if err := local.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	check, err := local.Publisher().One(ctx, tx)
	if err != nil {
		t.Fatal(err)
	}

	if check.ID != foreign.ID {
		t.Errorf("want: %v, got %v", foreign.ID, check.ID)
	}

	ranAfterSelectHook := false
	AddPublisherHook(boil.AfterSelectHook, func(ctx context.Context, e boil.ContextExecutor, o *Publisher) error {
		ranAfterSelectHook = true
		return nil
	})

	slice := GameSlice{&local}
	if err = local.L.LoadPublisher(ctx, tx, false, (*[]*Game)(&slice), nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Publisher == nil {
		t.Error("struct should have been eager loaded")
	}

	local.R.Publisher = nil
	if err = local.L.LoadPublisher(ctx, tx, true, &local, nil); err != nil {
		t.Fatal(err)
	}
	if local.R.Publisher == nil {
		t.Error("struct should have been eager loaded")
	}

	if !ranAfterSelectHook {
		t.Error("failed to run AfterSelect hook for relationship")
	}
}

func testGameToOneSetOpUserUsingOwnerUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Game
	var b, c User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, gameDBTypes, false, strmangle.SetComplement(gamePrimaryKeyColumns, gameColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*User{&b, &c} {
		err = a.SetOwnerUser(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.OwnerUser != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.OwnerUserGames[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if !queries.Equal(a.OwnerUserID, x.ID) {
			t.Error("foreign key was wrong value", a.OwnerUserID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.OwnerUserID))
		reflect.Indirect(reflect.ValueOf(&a.OwnerUserID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if !queries.Equal(a.OwnerUserID, x.ID) {
			t.Error("foreign key was wrong value", a.OwnerUserID, x.ID)
		}
	}
}

func testGameToOneRemoveOpUserUsingOwnerUser(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Game
	var b User

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, gameDBTypes, false, strmangle.SetComplement(gamePrimaryKeyColumns, gameColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, userDBTypes, false, strmangle.SetComplement(userPrimaryKeyColumns, userColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err = a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	if err = a.SetOwnerUser(ctx, tx, true, &b); err != nil {
		t.Fatal(err)
	}

	if err = a.RemoveOwnerUser(ctx, tx, &b); err != nil {
		t.Error("failed to remove relationship")
	}

	count, err := a.OwnerUser().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 0 {
		t.Error("want no relationships remaining")
	}

	if a.R.OwnerUser != nil {
		t.Error("R struct entry should be nil")
	}

	if !queries.IsValuerNil(a.OwnerUserID) {
		t.Error("foreign key value should be nil")
	}

	if len(b.R.OwnerUserGames) != 0 {
		t.Error("failed to remove a from b's relationships")
	}
}

func testGameToOneSetOpPublisherUsingPublisher(t *testing.T) {
	var err error

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()

	var a Game
	var b, c Publisher

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, &a, gameDBTypes, false, strmangle.SetComplement(gamePrimaryKeyColumns, gameColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &b, publisherDBTypes, false, strmangle.SetComplement(publisherPrimaryKeyColumns, publisherColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}
	if err = randomize.Struct(seed, &c, publisherDBTypes, false, strmangle.SetComplement(publisherPrimaryKeyColumns, publisherColumnsWithoutDefault)...); err != nil {
		t.Fatal(err)
	}

	if err := a.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}
	if err = b.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Fatal(err)
	}

	for i, x := range []*Publisher{&b, &c} {
		err = a.SetPublisher(ctx, tx, i != 0, x)
		if err != nil {
			t.Fatal(err)
		}

		if a.R.Publisher != x {
			t.Error("relationship struct not set to correct value")
		}

		if x.R.Games[0] != &a {
			t.Error("failed to append to foreign relationship struct")
		}
		if a.PublisherID != x.ID {
			t.Error("foreign key was wrong value", a.PublisherID)
		}

		zero := reflect.Zero(reflect.TypeOf(a.PublisherID))
		reflect.Indirect(reflect.ValueOf(&a.PublisherID)).Set(zero)

		if err = a.Reload(ctx, tx); err != nil {
			t.Fatal("failed to reload", err)
		}

		if a.PublisherID != x.ID {
			t.Error("foreign key was wrong value", a.PublisherID, x.ID)
		}
	}
}

func testGamesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Game{}
	if err = randomize.Struct(seed, o, gameDBTypes, true, gameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Game struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testGamesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Game{}
	if err = randomize.Struct(seed, o, gameDBTypes, true, gameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Game struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := GameSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testGamesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &Game{}
	if err = randomize.Struct(seed, o, gameDBTypes, true, gameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Game struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := Games().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	gameDBTypes = map[string]string{`ID`: `uuid`, `Name`: `character varying`, `PublisherID`: `uuid`, `L2TX`: `character varying`, `OwnerUserID`: `uuid`}
	_           = bytes.MinRead
)

func testGamesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(gamePrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(gameAllColumns) == len(gamePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Game{}
	if err = randomize.Struct(seed, o, gameDBTypes, true, gameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Game struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Games().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, gameDBTypes, true, gamePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Game struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testGamesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(gameAllColumns) == len(gamePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &Game{}
	if err = randomize.Struct(seed, o, gameDBTypes, true, gameColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Game struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := Games().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, gameDBTypes, true, gamePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Game struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(gameAllColumns, gamePrimaryKeyColumns) {
		fields = gameAllColumns
	} else {
		fields = strmangle.SetComplement(
			gameAllColumns,
			gamePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := GameSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testGamesUpsert(t *testing.T) {
	t.Parallel()

	if len(gameAllColumns) == len(gamePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := Game{}
	if err = randomize.Struct(seed, &o, gameDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Game struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Game: %s", err)
	}

	count, err := Games().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, gameDBTypes, false, gamePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Game struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert Game: %s", err)
	}

	count, err = Games().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}