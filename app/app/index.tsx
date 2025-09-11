import { Text, View } from '@/components/Themed';
import { Image, StyleSheet, TextInput, KeyboardAvoidingView, Platform, Button, Pressable } from 'react-native';
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
              placeholderTextColor='#7B8794'
            />
             <Pressable 
               style={({ pressed }) => [
                 styles.niceButton,
                 { opacity: pressed ? 0.8 : 1 }
               ]} 
               onPress={() => {}}
             >
               <Text style={styles.niceButtonText}>Search</Text>
             </Pressable>
             <Pressable 
               style={({ pressed }) => [
                 styles.niceButtonSecondary,
                 { opacity: pressed ? 0.8 : 1 }
               ]} 
               onPress={() => {}}
             >
               <Text style={styles.niceButtonTextSecondary}>Saved Restaurants</Text>
             </Pressable>
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
    borderColor: '#FF6B35',
    color: '#192A56',
    borderWidth: 2,
    borderRadius: 12,
    paddingHorizontal: 16,
    fontSize: 16,
    backgroundColor: 'white',
  },
  niceButton: {
    width: '90%',
    height: 52,
    backgroundColor: '#FF6B35',
    borderRadius: 16,
    justifyContent: 'center',
    alignItems: 'center',
    marginTop: 20,
    shadowColor: '#FF6B35',
    shadowOffset: { width: 0, height: 4 },
    shadowOpacity: 0.4,
    shadowRadius: 8,
    elevation: 8,
  },
  niceButtonText: {
    color: 'white',
    fontSize: 18,
    fontWeight: '600',
  },
  niceButtonSecondary: {
    width: '90%',
    height: 52,
    backgroundColor: 'rgba(255,255,255,0.1)',
    borderColor: 'white',
    borderWidth: 2,
    borderRadius: 16,
    justifyContent: 'center',
    alignItems: 'center',
    marginTop: 12,
  },
  niceButtonTextSecondary: {
    color: 'white',
    fontSize: 18,
    fontWeight: '600',
  },
});
