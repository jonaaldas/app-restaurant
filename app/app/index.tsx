import { Text, View } from '@/components/Themed';
import { Image, StyleSheet, TextInput, KeyboardAvoidingView, Platform } from 'react-native';
import { SafeAreaView } from 'react-native-safe-area-context';

export default function Index() {
  return (
    <View style={styles.container}>
      <Image resizeMode='cover' source={require('@/assets/images/res.jpg')} style={styles.image} />
      <View style={styles.overlay} />
      <KeyboardAvoidingView 
        behavior={Platform.OS === 'ios' ? 'padding' : 'height'}
        style={styles.keyboardView}
      >
        <SafeAreaView style={styles.safeArea}>
          <View style={styles.content} lightColor="transparent" darkColor="transparent">
            <Text style={styles.title}>Discover Amazing Restaurants</Text>
            <Text style={styles.description}>Find your next favorite dining experience in your neighborhood</Text>
            <TextInput 
              style={styles.input} 
              placeholder='Search for a restaurant' 
              placeholderTextColor='black'
            />
          </View>
        </SafeAreaView>
      </KeyboardAvoidingView>
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    position: 'relative',
  },
  keyboardView: {
    flex: 1,
  },
  safeArea: {
    flex: 1,
    justifyContent: 'center',
    alignItems: 'center',
  },
  content: {
    display: 'flex',
    width: '100%',
    flexDirection: 'column',
    justifyContent: 'center',
    alignItems: 'center',
    paddingVertical: 32,
    paddingHorizontal: 20,
  },
  image: {
    height: '100%',
    width: '100%',
    position: 'absolute',
    top: 0,
    left: 0,
    right: 0,
    bottom: 0,
  },
  overlay: {
    height: '100%',
    width: '100%',
    position: 'absolute',
    top: 0,
    left: 0,
    right: 0,
    bottom: 0,
    backgroundColor: 'rgba(0,0,0,0.4)',
  },
  title: {
    fontSize: 32,
    fontWeight: 'bold',
    color: 'white',
    textAlign: 'center',
    marginBottom: 8,
  },
  description: {
    textAlign: 'center',
    fontSize: 18,
    color: 'white',
    marginBottom: 24,
    lineHeight: 24,
  },
  input: {
    width: '90%',
    height: 48,
    borderColor: 'white',
    color: 'black',
    borderWidth: 1,
    borderRadius: 12,
    paddingHorizontal: 16,
    fontSize: 16,
    backgroundColor: 'white',
  },
});
