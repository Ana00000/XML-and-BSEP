<template>
  <div class="searchDiv">

        <div class="spacingOne" />
        <div class="title">
          <h1>You can find locations here!</h1>
          <div class="welcoming"> Search locations</div>
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
              v-model="selectedLocation"
              single
            >
              <template v-for="(location, id) in locations">
                <v-list-item
                  :key="location.id"
                  :value="location"
                  v-on:click="redirectToSelectedLocation"
                >
                  <template>
                    <v-list-item-content>
                      <v-list-item-subtitle
                        v-text="'LOCATION CITY: ' + location.country"
                        class="containerDiv"
                      ></v-list-item-subtitle>
                       <v-list-item-subtitle
                        v-text="'LOCATION COUNTRY: ' + location.city"
                        class="containerDiv"
                      ></v-list-item-subtitle>
                       <v-list-item-subtitle
                        v-text="'LOCATION STREET NAME: ' + location.streetName"
                        class="containerDiv"
                      ></v-list-item-subtitle>
                       <v-list-item-subtitle
                        v-text="'LOCATION STREET NUMBER: ' + location.streetNumber"
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
                  v-if="`A-${id}` < locations.length - 1"
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
  name: "SearchLocationsNotRegistered",
  data: () => ({
    searchInput: "",
    locations: [],
    locationsIds: [],
    locationsCopy: [],
    selectedLocation: null,
    locationString: "",
  }),
  mounted() {
    this.init();
  },
  methods: {
    init() {
      this.$http
        .get("https://localhost:8080/api/post/find_all_locations_for_public_posts/")
        .then((resp) => {
          //this.locations = resp.data;
          this.locations = resp.data;
          this.locationsCopy = resp.data;
        })
        .catch(console.log);
    },

    searchQuery() {
      var resultOfSearch = [];
      for (var i = 0; i < this.locationsCopy.length; i++) {
        if (this.locationsCopy[i].city.toLowerCase().includes(this.searchInput.toLowerCase()) || this.locationsCopy[i].country.toLowerCase().includes(this.searchInput.toLowerCase()) || this.locationsCopy[i].streetName.toLowerCase().includes(this.searchInput.toLowerCase()))
          resultOfSearch.push(this.locationsCopy[i]);
      }
      this.locations = resultOfSearch;
    },
    redirectToSelectedLocation() {

      this.locationString = this.selectedLocation.country+","+this.selectedLocation.city+","+this.selectedLocation.streetName+","+this.selectedLocation.streetNumber
      console.log("STRING: ")
      console.log(this.locationString)
      localStorage.setItem("selectedLocationString", this.locationString);
      localStorage.setItem("selectedLocationId", this.selectedLocation.id);
      window.location.href =
        "https://localhost:8081/postsForSelectedLocationNotRegistered";
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
.allLocations {
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