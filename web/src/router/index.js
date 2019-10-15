import Vue from 'vue'
import Router from 'vue-router'
import Main from '@/view/main/main'
import Generator from '@/view/generator'

Vue.use(Router);

const router = new Router({
  routes: [{
    path: '/',
    name: 'Main',
    component: Main,
    children: [{
      path: '/generator',
      name: 'Generator',
      component: Generator
    }]
  }]
});


router.beforeEach((to, from, next) => {
  if (to.path === '/') {
    router.push({
      name: 'Generator'
    })
  } else {
    next()
  }
});

export default router
