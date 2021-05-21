<template>
  <div class="searchDiv">

        <div class="spacingOne" />
        <div class="title">
          <h1>You can find tags here!</h1>
          <div class="welcoming"> Search tags</div>
        </div>
        <div class="spacingTwo" />

    <v-container>
      <v-layout row wrap>
        <v-card
          class="mx-auto"
          style="width: 90%; height: 300px; overflow-y: scroll"
        >
          <v-toolbar color="light-blue darken-4">
            <v-text-field
              hide-details
              prepend-icon="mdi-magnify"
              single-line
              v-model="searchInput"
              v-on:keyup="searchQuery()"
            />
          </v-toolbar>
          <v-list two-line>
            <v-list-item-group
              active-class="indigo--text"
              v-model="selectedTag"
              single
            >
              <template v-for="(tag, id) in tags">
                <v-list-item
                  :key="tag.id"
                  :value="tag"
                  v-on:click="redirectToSelectedTag"
                >
                  <template>
                    <v-list-item-content>
                      <v-list-item-subtitle
                        v-text="'TAG NAME: ' + tag.name"
                        class="containerDiv"
                      ></v-list-item-subtitle>
                      <v-list-item-subtitle
                        v-text="' '"
                        class="emptyContentClass"
                      ></v-list-item-subtitle>
                    </v-list-item-content>
                  </template>
                </v-list-item>
                <v-divider
                  v-if="`A-${id}` < tags.length - 1"
                  :key="`A-${id}`"
                />
              </template>
            </v-list-item-group>
          </v-list>
        </v-card>
      </v-layout>
    </v-container>
  </div>
</template>

<script>
export default {
  name: "SearchTagsRegistered",
  data: () => ({
    searchInput: "",
    tags: [],
    tagsCopy: [],
    selectedTag: null,
  }),
  mounted() {
    this.init();
  },
  methods: {
    init() {
      this.$http
        .get("http://localhost:8084/find_all_tags_for_public_and_friends_posts?id="+localStorage.getItem("userId"))
        .then((resp) => {
          this.tags = resp.data;
          this.tagsCopy = resp.data;
        })
        .catch(console.log);
    },

    searchQuery() {
      var resultOfSearch = [];
      for (var i = 0; i < this.tagsCopy.length; i++) {
        if (
          this.tagsCopy[i].name
            .toLowerCase()
            .includes(this.searchInput.toLowerCase())
        )
          resultOfSearch.push(this.tagsCopy[i]);
      }
      this.tags = resultOfSearch;
    },
    redirectToSelectedTag() {
      localStorage.setItem("selectedTagName", this.selectedTag.name);
      localStorage.setItem("selectedTagId", this.selectedTag.id);
      window.location.href =
        "http://localhost:8081/selectedTagForRegistered";
    },
  },
};
</script>

<style scoped>
.sort {
  padding-top: 15px;
  padding-bottom: 15px;
}

.cardClass {
  display: none; 
}
.template {
    min-height: 1000px;
}
.allTags {
    position: absolute;
    right: 500px;
    top: 490px;
}
.welcoming {
    font-weight: bolder;
    font-size: 25px;
    margin-left: 7%;
}
.searchDiv {
     height: 840px; 
    
}
.containerDiv{
  font-weight: bolder;
  font-size: 20px;
}

.spacingOne {
  height: 50px;
}

.title {
  margin-left: 40%;
}

.spacingTwo {
  height: 100px;
}
</style>