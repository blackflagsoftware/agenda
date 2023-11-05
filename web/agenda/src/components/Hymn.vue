<script setup>
import axios from "axios";
</script>

<template>
	<v-select :label="label" variant="outlined" v-model="hymnNumber" density="compact" :items="items" item-title="name" item-value="id" v-on:update:modelValue="onChange"></v-select>
</template>

<script>
export default {
	name: "Hymn",
	props: [
		"hymnNumberIn",
		"hymns",
		"label"
	],
	data() {
		return {
			hymnNumber: this.hymnNumberIn,
			items: this.hymns,
		}
	},
	methods: {
		onChange: function() {
			this.$emit("hymn-number-out", this.hymnNumber.toString())
		}
	},
	watch: {
		hymnNumberIn: {
			handler(newItem, oldItem) {
				this.hymnNumber = Number(newItem)
			},
			immediate: true
		}
	}
}
</script>