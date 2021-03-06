import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/logIn',
    name: 'LogIn',
    component: () => import('../views/LogIn.vue')
  },
  {
    path: '/register',
    name: 'Register',
    component: () => import('../views/Register.vue')
  },
  {
    path: '/confirmRegistration/:confirmationToken/:userId',
    name: 'ConfirmRegistration',
    component: () => import('../views/ConfirmRegistration.vue')
  },
  {
    path: '/recoverPasswordEmail',
    name: 'RecoverPasswordEmail',
    component: () => import('../views/RecoverPasswordEmail.vue')
  },
  {
    path: '/changePasswordByToken/:token/:id',
    name: 'ChangePasswordByToken',
    component: () => import('../views/ChangePasswordByToken.vue')
  },
  {
    path: '/updateProfile',
    name: 'UpdateProfile',
    component: () => import('../views/UpdateProfile.vue')
  },
  {
    path: '/myMedia',
    name: 'MyMedia',
    component: () => import('../views/MyMedia.vue')
  },
  {
    path: '/searchUsers',
    name: 'SearchUsers',
    component: () => import('../views/SearchUsers.vue')
  },
  {
    path: '/searchUsersForNotRegistered',
    name: 'SearchUsersForNotRegistered',
    component: () => import('../views/SearchUsersForNotRegistered.vue')
  },
  {
    path: '/selectedUser',
    name: 'SelectedUser',
    component: () => import('../views/SelectedUser.vue')
  },
  {
    path: '/selectedUserForNotRegistered',
    name: 'SelectedUserForNotRegistered',
    component: () => import('../views/SelectedUserForNotRegistered.vue')
  },
  {
    path: '/createPost',
    name: 'CreatePost',
    component: () => import('../views/CreatePost.vue')
  },
  {
    path: '/createStory',
    name: 'CreateStory',
    component: () => import('../views/CreateStory.vue')
  },
  {
    path: '/createPostAlbum',
    name: 'CreatePostAlbum',
    component: () => import('../views/CreatePostAlbum.vue')
  },
  {
    path: '/createStoryAlbum',
    name: 'CreateStoryAlbum',
    component: () => import('../views/CreateStoryAlbum.vue')
  },
  {
    path: '/myCollections',
    name: 'MyCollections',
    component: () => import('../views/MyCollections.vue')
  },
  {
    path: '/friendsPosts',
    name: 'FriendsPosts',
    component: () => import('../views/FriendsPosts.vue')
  },
  {
    path: '/publicPostsAll',
    name: 'PublicPostsAll',
    component: () => import('../views/PublicPostsAll.vue')
  },
  {
    path: '/publicPostsAllForNotRegistered',
    name: 'PublicPostsAllForNotRegistered',
    component: () => import('../views/PublicPostsAllForNotRegistered.vue')
  },
  {
    path: '/friendsStories',
    name: 'FriendsStories',
    component: () => import('../views/FriendsStories.vue')
  },
  {
    path: '/publicStoriesAll',
    name: 'PublicStoriesAll',
    component: () => import('../views/PublicStoriesAll.vue')
  },
  {
    path: '/publicStoriesAllForNotRegistered',
    name: 'PublicStoriesAllForNotRegistered',
    component: () => import('../views/PublicStoriesAllForNotRegistered.vue')
  },
  {
    path: '/reactedPosts',
    name: 'ReactedPosts',
    component: () => import('../views/ReactedPosts.vue')
  },
  {
    path: '/followRequests',
    name: 'FollowRequests',
    component: () => import('../views/FollowRequests.vue')
  },
  {
    path: '/selectedFollowRequest',
    name: 'SelectedFollowRequest',
    component: () => import('../views/SelectedFollowRequest.vue')
  },
  {
    path: '/postById',
    name: 'PostById',
    component: () => import('../views/PostById.vue')
  },
  {
    path: '/postByIdWithoutActivity',
    name: 'PostByIdWithoutActivity',
    component: () => import('../views/PostByIdWithoutActivity.vue')
  },
  {
    path: '/storyByIdWithoutActivity',
    name: 'StoryByIdWithoutActivity',
    component: () => import('../views/StoryByIdWithoutActivity.vue')
  },
  {
    path: '/postAlbumByIdWithoutActivity',
    name: 'PostAlbumByIdWithoutActivity',
    component: () => import('../views/PostAlbumByIdWithoutActivity.vue')
  },
  {
    path: '/storyAlbumByIdWithoutActivity',
    name: 'StoryAlbumByIdWithoutActivity',
    component: () => import('../views/StoryAlbumByIdWithoutActivity.vue')
  },
  {
    path: '/addCloseFriends',
    name: 'AddCloseFriends',
    component: () => import('../views/AddCloseFriends.vue')
  },
  {
    path: '/storiesOfStoryHighlight',
    name: 'StoriesOfStoryHighlight',
    component: () => import('../views/StoriesOfStoryHighlight.vue')
  },
  {
    path: '/searchTagsNotRegistered',
    name: 'SearchTagsNotRegistered',
    component: () => import('../views/SearchTagsNotRegistered.vue')
  },
  {
    path: '/searchTagsRegistered',
    name: 'SearchTagsRegistered',
    component: () => import('../views/SearchTagsRegistered.vue')
  },
  {
    path: '/searchLocationsNotRegistered',
    name: 'SearchLocationsNotRegistered',
    component: () => import('../views/SearchLocationsNotRegistered.vue')
  },
  {
    path: '/searchLocationsRegistered',
    name: 'SearchLocationsRegistered',
    component: () => import('../views/SearchLocationsRegistered.vue')
  },
  {
    path: '/postsForSelectedTagNotRegistered',
    name: 'PostsForSelectedTagNotRegistered',
    component: () => import('../views/PostsForSelectedTagNotRegistered.vue')
  },
  {
    path: '/postsForSelectedTagRegistered',
    name: 'PostsForSelectedTagRegistered',
    component: () => import('../views/PostsForSelectedTagRegistered.vue')
  },
  {
    path: '/postsForSelectedLocationNotRegistered',
    name: 'PostsForSelectedLocationNotRegistered',
    component: () => import('../views/PostsForSelectedLocationNotRegistered.vue')
  },
  {
    path: '/postsForSelectedLocationRegistered',
    name: 'PostsForSelectedLocationRegistered',
    component: () => import('../views/PostsForSelectedLocationRegistered.vue')
  },
  {
    path: '/unauthorizedPage',
    name: 'UnauthorizedPage',
    component: () => import('../views/UnauthorizedPage.vue')
  },
  {
    path: '/forbiddenPage',
    name: '/ForbiddenPage',
    component: () => import('../views/ForbiddenPage.vue')
  },
  {
    path: '/verificationRequests',
    name: 'VerificationRequests',
    component: () => import('../views/VerificationRequests.vue')
  },
  {
    path: '/selectedVerificationRequest',
    name: 'SelectedVerificationRequest',
    component: () => import('../views/SelectedVerificationRequest.vue')
  },
  {
    path: '/createVerificationRequest',
    name: 'CreateVerificationRequest',
    component: () => import('../views/CreateVerificationRequest.vue')
  },
  {
    path: '/updateSettings',
    name: '/UpdateSettings',
    component: () => import('../views/UpdateSettings.vue')
  },
  {
    path: '/updateNotifications',
    name: '/UpdateNotifications',
    component: () => import('../views/UpdateNotifications.vue')
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.VUE_APP_BASE_URL,
  routes
})

export default router
