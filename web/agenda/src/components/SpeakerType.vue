<template>
	<v-select label="Type" variant="outlined" density="compact" v-model="value" :items="['Speaker', 'Hymn', 'Muscial Number']" v-on:update:modelValue="onChangeValue"></v-select>
</template>

<script>
export default {
	name: "SpeakerType",
	data() {
		return {
			value: ""
		}
	},
	methods: {
		onChangeValue: function() {
			this.$emit("speaker-type-out", this.value)
		}
	}
}
</script>