<template>
  <div>
    <img :src="require(`../../../Media/${img}`)" alt class="icon" />
  </div>
</template>

<script>
export default {
  name: "Posts",
  data: () => ({
    img: "026579356Screenshot (62).png"
  }),
};
</script>
