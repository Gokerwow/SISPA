import { createRouter, createWebHistory } from 'vue-router'
import DashboardLayout from '../layouts/dashboardLayout.vue'
import CustomerView from '../views/customers/customerView.vue'
import CustomerCreate from '@/views/customers/customerCreate.vue'
import CustomerEdit from '@/views/customers/customerEdit.vue'
import CustomerDetail from '@/views/customers/customerDetail.vue'
import ServiceView from '@/views/service/serviceView.vue'
import ServiceCreate from '@/views/service/serviceCreate.vue'
import ServiceEdit from '@/views/service/serviceEdit.vue'
import TherapistView from '@/views/therapist/therapistView.vue'
import TherapistCreate from '@/views/therapist/therapistCreate.vue'
import TherapistEdit from '@/views/therapist/therapistEdit.vue'
import TransactionView from '@/views/transaction/transactionView.vue'
import TransactionCreate from '@/views/transaction/transactionCreate.vue'
import TransactionDetail from '@/views/transaction/transactionDetail.vue'
import TransactionEdit from '@/views/transaction/transactionEdit.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      component: DashboardLayout, // The Shell
      children: [
        {
          path: '', // Default route (Dashboard Home)
          redirect: '/customers' // For now, just send us to customers
        },
        {
          path: 'customers', // URL: /customers
          name: 'customers',
          children: [
            {
              path: '', // URL: /customers
              name: 'Customers',
              component: CustomerView // The Module
            },
            {
              path: ':id', // URL: /customers
              name: 'Customer Detail',
              component: CustomerDetail // The Module
            },
            {
              path: 'create', // URL: /customers
              name: 'Create Customers',
              component: CustomerCreate // The Module
            },
            {
              path: 'edit/:id', // The colon (:) makes 'id' a variable
              name: 'Edit Customer',
              component: CustomerEdit
            }
          ]
        },
        {
          path: 'services',
          name: 'services',
          children: [
            {
              path: '',
              name: 'Services',
              component: ServiceView
            },
            {
              path: 'create',
              name: 'Create Services',
              component: ServiceCreate
            },
            {
              path: 'edit/:id',
              name: 'Edit Services',
              component: ServiceEdit
            },
          ]
        },
        {
          path: 'therapists',
          name: 'therapists',
          children: [
            {
              path: '',
              name: 'Therapists',
              component: TherapistView
            },
            {
              path: 'create',
              name: 'Create Therapists',
              component: TherapistCreate
            },
            {
              path: 'edit/:id',
              name: 'Edit Therapists',
              component: TherapistEdit
            },
          ]
        },
        {
          path: 'transactions',
          name: 'transactions',
          children: [
            {
              path: '',
              name: 'Transactions',
              component: TransactionView
            },
            {
              path: 'create',
              name: 'Create Transactions',
              component: TransactionCreate
            },
            {
              path: 'edit/:id',
              name: 'Edit Transactions',
              component: TransactionEdit
            },
            {
              path: ':id',
              name: 'Transactions Detail',
              component: TransactionDetail
            },
          ]
        },
      ]
    }
  ]
})

export default router
