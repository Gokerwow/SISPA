<script setup lang="ts">
import { reactive, watch } from 'vue'
import {
  UserCircle,
  Phone,
  MapPin,
  Save,
  X,
  Calendar,
  User
} from 'lucide-vue-next'
import type { Therapist } from '@/types/therapist';

// Form data based on Therapist interface
const formData = reactive({
  id: 0,
  name: '',
  phone: '',
  address: '',
  commission_rate: 0,
  joined_date: '',
  exit_date: null,
  status: '',
  created_at: '',
  updated_at: '',
})

const props = defineProps<{
  initialData?: Therapist,
  isSubmitting?: boolean
}>()

const emit = defineEmits<{
  (e: 'submit', data: Therapist): void
}>()

watch(() => props.initialData, (newData) => {
  if (newData) {
    Object.assign(formData, newData)
  }
}, { immediate: true })

// Submit form
const handleSubmit = () => {
  emit('submit', {...formData})
}

// Status options
const statusOptions = [
  { value: 'active', label: 'Active', color: 'emerald' },
  { value: 'inactive', label: 'Inactive', color: 'slate' },
  { value: 'on leave', label: 'On Leave', color: 'amber' },
  { value: 'terminated', label: 'Terminated', color: 'red' }
]

// Get status color classes
const getStatusColorClasses = (color: string, isSelected: boolean) => {
  const colors: Record<string, { border: string, bg: string }> = {
    emerald: {
      border: isSelected ? 'border-emerald-300 bg-emerald-50' : 'border-slate-200 bg-white hover:border-slate-300',
      bg: 'bg-emerald-200 text-emerald-700'
    },
    slate: {
      border: isSelected ? 'border-slate-300 bg-slate-50' : 'border-slate-200 bg-white hover:border-slate-300',
      bg: 'bg-slate-200 text-slate-600'
    },
    amber: {
      border: isSelected ? 'border-amber-300 bg-amber-50' : 'border-slate-200 bg-white hover:border-slate-300',
      bg: 'bg-amber-200 text-amber-700'
    },
    red: {
      border: isSelected ? 'border-red-300 bg-red-50' : 'border-slate-200 bg-white hover:border-slate-300',
      bg: 'bg-red-200 text-red-700'
    }
  }
  return colors[color]
}
</script>

<template>
  <div class="space-y-6">

    <!-- Header -->
    <div class="flex items-center justify-between">
      <div>
        <h1 class="text-3xl font-bold bg-gradient-to-r from-rose-600 via-purple-600 to-indigo-600 bg-clip-text text-transparent mb-2">
          Add New Therapist
        </h1>
        <p class="text-slate-500 text-sm">Create a new therapist profile for your spa team</p>
      </div>
      <button class="p-3 hover:bg-slate-100 rounded-xl transition-all text-slate-600">
        <X class="w-5 h-5" />
      </button>
    </div>

    <form @submit.prevent="handleSubmit" class="space-y-6">

      <!-- Personal Information -->
      <div class="bg-white/80 backdrop-blur-xl rounded-2xl p-6 border border-white/60 shadow-lg">
        <h2 class="text-lg font-semibold text-slate-800 mb-5 flex items-center gap-2">
          <User class="w-5 h-5 text-purple-500" />
          Personal Information
        </h2>

        <div class="space-y-5">
          <!-- Therapist Name -->
          <div>
            <label class="block text-sm font-medium text-slate-700 mb-2">
              Full Name <span class="text-red-500">*</span>
            </label>
            <div class="relative">
              <UserCircle class="absolute left-4 top-1/2 -translate-y-1/2 w-5 h-5 text-slate-400" />
              <input
                v-model="formData.name"
                type="text"
                required
                placeholder="Enter therapist full name"
                class="w-full pl-12 pr-4 py-3 bg-slate-50 border border-slate-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-purple-400 focus:bg-white transition-all"
              />
            </div>
          </div>

          <div class="grid grid-cols-1 md:grid-cols-2 gap-5">
            <!-- Phone -->
            <div>
              <label class="block text-sm font-medium text-slate-700 mb-2">
                Phone Number <span class="text-red-500">*</span>
              </label>
              <div class="relative">
                <Phone class="absolute left-4 top-1/2 -translate-y-1/2 w-5 h-5 text-slate-400" />
                <input
                  v-model="formData.phone"
                  type="tel"
                  required
                  placeholder="+62 812-3456-7890"
                  class="w-full pl-12 pr-4 py-3 bg-slate-50 border border-slate-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-purple-400 focus:bg-white transition-all"
                />
              </div>
            </div>

            <!-- Joined Date -->
            <div>
              <label class="block text-sm font-medium text-slate-700 mb-2">
                Joined Date
              </label>
              <div class="relative">
                <Calendar class="absolute left-4 top-1/2 -translate-y-1/2 w-5 h-5 text-slate-400" />
                <input
                  v-model="formData.joined_date"
                  type="date"
                  class="w-full pl-12 pr-4 py-3 bg-slate-50 border border-slate-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-purple-400 focus:bg-white transition-all"
                />
              </div>
            </div>
          </div>

          <!-- Address -->
          <div>
            <label class="block text-sm font-medium text-slate-700 mb-2">
              Address <span class="text-red-500">*</span>
            </label>
            <div class="relative">
              <MapPin class="absolute left-4 top-4 w-5 h-5 text-slate-400" />
              <textarea
                v-model="formData.address"
                rows="3"
                required
                placeholder="Enter full address"
                class="w-full pl-12 pr-4 py-3 bg-slate-50 border border-slate-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-purple-400 focus:bg-white transition-all resize-none"
              ></textarea>
            </div>
          </div>
        </div>
      </div>

      <!-- TODO: commision or payment for the therapist soon -->
      <!-- Commission Rate -->
      <!-- <div class="bg-white/80 backdrop-blur-xl rounded-2xl p-6 border border-white/60 shadow-lg">
        <h2 class="text-lg font-semibold text-slate-800 mb-5 flex items-center gap-2">
          <Percent class="w-5 h-5 text-purple-500" />
          Commission Rate
        </h2>

        <div class="bg-gradient-to-br from-purple-50 to-indigo-50 border-2 border-purple-200 rounded-xl p-5">
          <label class="block text-sm font-semibold text-slate-800 mb-3">
            Commission Percentage <span class="text-red-500">*</span>
          </label>
          <div class="relative">
            <Percent class="absolute left-4 top-1/2 -translate-y-1/2 w-5 h-5 text-slate-500" />
            <input
              v-model.number="formData.commission_rate"
              type="number"
              required
              min="0"
              max="100"
              step="1"
              placeholder="0"
              class="w-full pl-12 pr-16 py-3 bg-white border-2 border-purple-200 rounded-xl focus:outline-none focus:ring-2 focus:ring-purple-400 transition-all font-semibold text-slate-800 text-lg"
            />
            <span class="absolute right-4 top-1/2 -translate-y-1/2 text-slate-500 font-medium">%</span>
          </div>
          <p class="text-xs text-slate-600 mt-2">Enter the commission rate (0-100%)</p>
        </div> -->

        <!-- Commission Preview -->
        <!-- <div class="mt-5 p-4 bg-gradient-to-br from-slate-50 to-purple-50 border border-purple-200 rounded-xl">
          <p class="text-xs text-slate-600 mb-2">Example Calculation:</p>
          <div class="grid grid-cols-2 gap-3 text-sm">
            <div class="p-3 bg-white rounded-lg border border-slate-200">
              <p class="text-xs text-slate-500 mb-1">Service Price</p>
              <p class="font-bold text-slate-800">Rp 100,000</p>
            </div>
            <div class="p-3 bg-purple-50 rounded-lg border border-purple-200">
              <p class="text-xs text-slate-500 mb-1">Therapist Earns</p>
              <p class="font-bold text-purple-700">Rp {{ (100000 * formData.commission_rate / 100).toLocaleString('id-ID') }}</p>
            </div>
          </div>
        </div>
      </div> -->

      <!-- Status -->
      <div class="bg-white/80 backdrop-blur-xl rounded-2xl p-6 border border-white/60 shadow-lg">
        <h2 class="text-lg font-semibold text-slate-800 mb-5 flex items-center gap-2">
          <UserCircle class="w-5 h-5 text-purple-500" />
          Employment Status
        </h2>

        <div class="grid grid-cols-1 md:grid-cols-2 gap-3">
          <label
            v-for="option in statusOptions"
            :key="option.value"
            :class="['flex items-center gap-4 p-5 rounded-xl border-2 cursor-pointer transition-all', getStatusColorClasses(option.color, formData.status === option.value).border]">
            <input
              v-model="formData.status"
              type="radio"
              :value="option.value"
              class="w-5 h-5 text-purple-500 focus:ring-purple-400"
            />
            <div class="flex-1">
              <p class="font-semibold text-slate-800 flex items-center gap-2">
                {{ option.label }}
                <span :class="['px-2 py-0.5 rounded-full text-xs font-bold', getStatusColorClasses(option.color, false).bg]">
                  {{ option.value }}
                </span>
              </p>
              <p class="text-xs text-slate-600 mt-1">
                <template v-if="option.value === 'active'">Available for appointments</template>
                <template v-else-if="option.value === 'inactive'">Not currently working</template>
                <template v-else-if="option.value === 'on leave'">Temporarily unavailable</template>
                <template v-else>No longer with the company</template>
              </p>
            </div>
          </label>
        </div>
      </div>

      <!-- Action Buttons -->
      <div class="flex items-center justify-end gap-4">
        <button
          type="button"
          class="px-6 py-3 bg-slate-100 hover:bg-slate-200 text-slate-700 rounded-2xl text-sm font-medium transition-all flex items-center gap-2"
        >
          <X class="w-4 h-4" />
          Cancel
        </button>

        <button
          type="submit"
          :disabled="isSubmitting"
          class="px-8 py-3 bg-gradient-to-r from-rose-500 via-purple-500 to-indigo-500 text-white rounded-2xl text-sm font-medium shadow-lg shadow-purple-300/40 hover:shadow-xl hover:shadow-purple-300/50 transition-all flex items-center gap-2 disabled:opacity-50 disabled:cursor-not-allowed"
        >
          <Save class="w-4 h-4" />
          <span v-if="!isSubmitting">Save Therapist</span>
          <span v-else>Saving...</span>
        </button>
      </div>

    </form>
  </div>
</template>

<style scoped>
/* Smooth transitions */
* {
  transition: all 0.2s cubic-bezier(0.4, 0, 0.2, 1);
}

/* Custom radio button */
input[type="radio"]:checked {
  background-color: rgb(168, 85, 247);
  border-color: rgb(168, 85, 247);
}

/* Remove number input arrows */
input[type="number"]::-webkit-inner-spin-button,
input[type="number"]::-webkit-outer-spin-button {
  -webkit-appearance: none;
  margin: 0;
}

input[type="number"] {
  -moz-appearance: textfield;
}

/* Smooth font rendering */
body {
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
}
</style>
