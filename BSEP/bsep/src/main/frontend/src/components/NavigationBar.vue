<template>
    <span>
        <v-navigation-drawer app temporary v-model="drawer" class="light-blue darken-4" dark disable-resize-watcher>
            <v-list>
                <template v-for="(item, index) in items">
                    <v-list-tile :key="index" >
                        <v-list-tile-content>
                            <v-card height="35px" class="ma-1 text-sm-button text-center">
                                <router-link class="router" :to="item.path">{{item.title}}</router-link>
                            </v-card>
                        </v-list-tile-content>
                    </v-list-tile>
                    <v-divider :key="`divider-${index}`"></v-divider>
                </template>
            </v-list>
        </v-navigation-drawer>
        <v-toolbar app dark class="light-blue darken-4">
            <v-app-bar-nav-icon class="hidden-md-and-up" @click="drawer = !drawer"></v-app-bar-nav-icon>
            <v-spacer class="hidden-md-and-up"></v-spacer>
            <v-btn flat class="hidden-sm-and-down ma-1" @click="drawer = !drawer">Menu</v-btn>
            <v-toolbar-title class="appTitlePos"><router-link class="appTitle" to="/">{{appTitle}}</router-link></v-toolbar-title>
            <v-spacer class="hidden-sm-and-down"></v-spacer>
            <div v-if="!isLogged">
                <v-btn flat class="hidden-sm-and-down">
                  <router-link class="router" to="/logIn">Log in</router-link>
                </v-btn>
            </div>
            <div v-else>
                <v-btn flat class="hidden-sm-and-down" v-on:click="logOff">Log off</v-btn>
            </div>
            <div v-if="!isLogged">
                <v-btn color="brown lighten-3" class="hidden-sm-and-down ma-1">
                  <router-link class="router" to="/register">Register</router-link>
                </v-btn>
            </div>
        </v-toolbar>
    </span>
</template>

<script>
export default {
    name: 'NavigationBar',
    computed:{
        isLogged: function(){
            var token = localStorage.getItem("token");
            return token !== "";
        }
    },
    data() {
        return {
            appTitle: 'Security system',
            drawer: false,
            isUserLogged:false,
            role : null,
            items: [
                { title: 'Home', path: '/' },
                { title: 'Certificates', path: '/certificates' },
                { title: 'New root certificate', path: '/newRootCertificate' },
                { title: 'Invalid certificates', path: '/invalidCertificates' },
                { title: 'Check certificate validity', path: '/checkCertificateValidity'}
            ]
        }
    },
    mounted() {
        this.init();
    },
    methods:{
    init(){
        this.role =localStorage.getItem('role');
        if (this.role == 'USER'){
            this.items = [
                { title: 'Home', path: '/' },
                { title: 'Certificates', path: '/certificates' },
                { title: 'Check certificate validity', path: '/checkCertificateValidity'}
            ]
        } else if (this.role == 'ADMIN') {
             [
                { title: 'Home', path: '/' },
                { title: 'Certificates', path: '/certificates' },
                { title: 'New root certificate', path: '/newRootCertificate' },
                { title: 'Invalid certificates', path: '/invalidCertificates' },
                { title: 'Check certificate validity', path: '/checkCertificateValidity'}
            ]
        } else {
            this.items = [
                { title: 'Home', path: '/' }
            ]
        }
    },
      logOff() {
        localStorage.setItem("token","");
        localStorage.setItem("role","NONE");
        window.location.href = "http://localhost:8081/";
      }
    }
}
</script>

<style scoped>
.appTitlePos {
    margin-left: 40%;
}

.appTitle {
    font-weight: bolder;
    font-size: 1.5em;
    height: 50px;
    text-decoration: none;
    color: white;
}

.router {
    text-decoration: none;
    color: white;
}
</style>