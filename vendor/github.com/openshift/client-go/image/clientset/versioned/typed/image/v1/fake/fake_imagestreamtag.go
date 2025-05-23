// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	v1 "github.com/openshift/api/image/v1"
	imagev1 "github.com/openshift/client-go/image/clientset/versioned/typed/image/v1"
	gentype "k8s.io/client-go/gentype"
)

// fakeImageStreamTags implements ImageStreamTagInterface
type fakeImageStreamTags struct {
	*gentype.FakeClientWithList[*v1.ImageStreamTag, *v1.ImageStreamTagList]
	Fake *FakeImageV1
}

func newFakeImageStreamTags(fake *FakeImageV1, namespace string) imagev1.ImageStreamTagInterface {
	return &fakeImageStreamTags{
		gentype.NewFakeClientWithList[*v1.ImageStreamTag, *v1.ImageStreamTagList](
			fake.Fake,
			namespace,
			v1.SchemeGroupVersion.WithResource("imagestreamtags"),
			v1.SchemeGroupVersion.WithKind("ImageStreamTag"),
			func() *v1.ImageStreamTag { return &v1.ImageStreamTag{} },
			func() *v1.ImageStreamTagList { return &v1.ImageStreamTagList{} },
			func(dst, src *v1.ImageStreamTagList) { dst.ListMeta = src.ListMeta },
			func(list *v1.ImageStreamTagList) []*v1.ImageStreamTag { return gentype.ToPointerSlice(list.Items) },
			func(list *v1.ImageStreamTagList, items []*v1.ImageStreamTag) {
				list.Items = gentype.FromPointerSlice(items)
			},
		),
		fake,
	}
}
