package nautobot_test

import (
	"net/url"
	"path"
	"testing"

	"github.com/google/uuid"
	"github.com/neverbeencloser/gonautobot/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gopkg.in/h2non/gock.v1"
)

const (
	testGitRepositoryID = "a1b2c3d4-e5f6-7890-abcd-ef1234567890"
)

func TestClient_GitRepositoryGet(t *testing.T) {
	gock.New(testURL).Get("extras/git-repositories/" + testGitRepositoryID + "/").Reply(200).
		File(path.Join("fixtures", "extras", "gitrepository_200_1.json"))

	id, err := uuid.Parse(testGitRepositoryID)
	require.NoError(t, err)

	resp, err := testClient.Extras.GitRepositoryGet(id)
	require.NoError(t, err)
	assert.Equal(t, "my-config-repo", resp.Name)
	assert.Equal(t, id, resp.ID)
	assert.Equal(t, "https://github.com/example/config-repo.git", resp.RemoteURL)
	assert.Equal(t, "main", resp.Branch)
}

func TestClient_GitRepositoryFilter(t *testing.T) {
	gock.New(testURL).Get("extras/git-repositories/").Reply(200).
		File(path.Join("fixtures", "extras", "gitrepositories_200_1.json"))

	q := &url.Values{}
	q.Set("name", "my-config-repo")

	resp, err := testClient.Extras.GitRepositoryFilter(q)
	require.NoError(t, err)
	assert.Len(t, resp, 2)
	assert.Equal(t, "my-config-repo", resp[0].Name)
}

func TestClient_GitRepositoryAll(t *testing.T) {
	gock.New(testURL).Get("extras/git-repositories/").Reply(200).
		File(path.Join("fixtures", "extras", "gitrepositories_200_1.json"))

	resp, err := testClient.Extras.GitRepositoryAll()
	require.NoError(t, err)

	assert.Len(t, resp, 2)
	assert.Equal(t, "my-config-repo", resp[0].Name)
	assert.Equal(t, "job-templates", resp[1].Name)
}

func TestClient_GitRepositoryCreate(t *testing.T) {
	newRepo := types.NewGitRepository{
		Name:             "my-config-repo",
		RemoteURL:        "https://github.com/example/config-repo.git",
		Branch:           "main",
		ProvidedContents: []string{"extras.configcontext"},
	}

	gock.New(testURL).Post("extras/git-repositories/").
		JSON(newRepo).
		Reply(201).
		File(path.Join("fixtures", "extras", "gitrepository_create_201_1.json"))

	resp, err := testClient.Extras.GitRepositoryCreate(newRepo)
	require.NoError(t, err)
	assert.Equal(t, newRepo.Name, resp.Name)
	assert.Equal(t, newRepo.RemoteURL, resp.RemoteURL)
	assert.Equal(t, uuid.MustParse(testGitRepositoryID), resp.ID)
}

func TestClient_GitRepositoryUpdate(t *testing.T) {
	updateData := map[string]any{
		"branch": "develop",
	}

	gock.New(testURL).Patch("extras/git-repositories/" + testGitRepositoryID + "/").
		JSON(updateData).
		Reply(200).
		File(path.Join("fixtures", "extras", "gitrepository_update_200_1.json"))

	id, err := uuid.Parse(testGitRepositoryID)
	require.NoError(t, err)

	resp, err := testClient.Extras.GitRepositoryUpdate(id, updateData)
	require.NoError(t, err)

	assert.Equal(t, "develop", resp.Branch)
	assert.Equal(t, id, resp.ID)
}

func TestClient_GitRepositoryDelete(t *testing.T) {
	gock.New(testURL).Delete("extras/git-repositories/" + testGitRepositoryID + "/").Reply(204)

	id, err := uuid.Parse(testGitRepositoryID)
	require.NoError(t, err)

	err = testClient.Extras.GitRepositoryDelete(id)
	require.NoError(t, err)
	assert.True(t, gock.IsDone())
}
