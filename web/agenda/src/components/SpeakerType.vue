<template>
	<v-select label="Type" variant="outlined" density="compact" v-model="value" :items="['Speaker', 'Hymn', 'Musical Number']" v-on:update:modelValue="onChangeValue"></v-select>
</template>

<script>
export default {
	name: "SpeakerType",
	props: [
		"valueIn"
	],
	data() {
		return {
			value: ""
		}
	},
	methods: {
		onChangeValue: function() {
			this.$emit("speaker-type-out", this.value)
		}
	},
	watch: {
		valueIn: {
			handler (newVal, oldVal) {
				this.value = newVal
			},
			immediate: true
		}
	}
}
</script>