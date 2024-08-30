// Copyright (c) 2024 Nxpod GmbH. All rights reserved.
// Licensed under the GNU Affero General Public License (AGPL).
// See License.AGPL.txt in the project root for license information.

// Generated by the protocol buffer compiler.  DO NOT EDIT!
// NO CHECKED-IN PROTOBUF GENCODE
// source: nxpod/v1/editor.proto
// Protobuf Java Version: 4.27.2

package io.nxpod.publicapi.v1;

public final class Editor {
  private Editor() {}
  static {
    com.google.protobuf.RuntimeVersion.validateProtobufGencodeVersion(
      com.google.protobuf.RuntimeVersion.RuntimeDomain.PUBLIC,
      /* major= */ 4,
      /* minor= */ 27,
      /* patch= */ 2,
      /* suffix= */ "",
      Editor.class.getName());
  }
  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistryLite registry) {
  }

  public static void registerAllExtensions(
      com.google.protobuf.ExtensionRegistry registry) {
    registerAllExtensions(
        (com.google.protobuf.ExtensionRegistryLite) registry);
  }
  public interface EditorReferenceOrBuilder extends
      // @@protoc_insertion_point(interface_extends:nxpod.v1.EditorReference)
      com.google.protobuf.MessageOrBuilder {

    /**
     * <code>string name = 1 [json_name = "name"];</code>
     * @return The name.
     */
    java.lang.String getName();
    /**
     * <code>string name = 1 [json_name = "name"];</code>
     * @return The bytes for name.
     */
    com.google.protobuf.ByteString
        getNameBytes();

    /**
     * <code>string version = 2 [json_name = "version"];</code>
     * @return The version.
     */
    java.lang.String getVersion();
    /**
     * <code>string version = 2 [json_name = "version"];</code>
     * @return The bytes for version.
     */
    com.google.protobuf.ByteString
        getVersionBytes();

    /**
     * <pre>
     * prefer_toolbox indicates whether the editor should be launched with the
     * JetBrains Toolbox instead of JetBrains Gateway
     * </pre>
     *
     * <code>bool prefer_toolbox = 3 [json_name = "preferToolbox"];</code>
     * @return The preferToolbox.
     */
    boolean getPreferToolbox();
  }
  /**
   * Protobuf type {@code nxpod.v1.EditorReference}
   */
  public static final class EditorReference extends
      com.google.protobuf.GeneratedMessage implements
      // @@protoc_insertion_point(message_implements:nxpod.v1.EditorReference)
      EditorReferenceOrBuilder {
  private static final long serialVersionUID = 0L;
    static {
      com.google.protobuf.RuntimeVersion.validateProtobufGencodeVersion(
        com.google.protobuf.RuntimeVersion.RuntimeDomain.PUBLIC,
        /* major= */ 4,
        /* minor= */ 27,
        /* patch= */ 2,
        /* suffix= */ "",
        EditorReference.class.getName());
    }
    // Use EditorReference.newBuilder() to construct.
    private EditorReference(com.google.protobuf.GeneratedMessage.Builder<?> builder) {
      super(builder);
    }
    private EditorReference() {
      name_ = "";
      version_ = "";
    }

    public static final com.google.protobuf.Descriptors.Descriptor
        getDescriptor() {
      return io.nxpod.publicapi.v1.Editor.internal_static_nxpod_v1_EditorReference_descriptor;
    }

    @java.lang.Override
    protected com.google.protobuf.GeneratedMessage.FieldAccessorTable
        internalGetFieldAccessorTable() {
      return io.nxpod.publicapi.v1.Editor.internal_static_nxpod_v1_EditorReference_fieldAccessorTable
          .ensureFieldAccessorsInitialized(
              io.nxpod.publicapi.v1.Editor.EditorReference.class, io.nxpod.publicapi.v1.Editor.EditorReference.Builder.class);
    }

    public static final int NAME_FIELD_NUMBER = 1;
    @SuppressWarnings("serial")
    private volatile java.lang.Object name_ = "";
    /**
     * <code>string name = 1 [json_name = "name"];</code>
     * @return The name.
     */
    @java.lang.Override
    public java.lang.String getName() {
      java.lang.Object ref = name_;
      if (ref instanceof java.lang.String) {
        return (java.lang.String) ref;
      } else {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        name_ = s;
        return s;
      }
    }
    /**
     * <code>string name = 1 [json_name = "name"];</code>
     * @return The bytes for name.
     */
    @java.lang.Override
    public com.google.protobuf.ByteString
        getNameBytes() {
      java.lang.Object ref = name_;
      if (ref instanceof java.lang.String) {
        com.google.protobuf.ByteString b =
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        name_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }

    public static final int VERSION_FIELD_NUMBER = 2;
    @SuppressWarnings("serial")
    private volatile java.lang.Object version_ = "";
    /**
     * <code>string version = 2 [json_name = "version"];</code>
     * @return The version.
     */
    @java.lang.Override
    public java.lang.String getVersion() {
      java.lang.Object ref = version_;
      if (ref instanceof java.lang.String) {
        return (java.lang.String) ref;
      } else {
        com.google.protobuf.ByteString bs =
            (com.google.protobuf.ByteString) ref;
        java.lang.String s = bs.toStringUtf8();
        version_ = s;
        return s;
      }
    }
    /**
     * <code>string version = 2 [json_name = "version"];</code>
     * @return The bytes for version.
     */
    @java.lang.Override
    public com.google.protobuf.ByteString
        getVersionBytes() {
      java.lang.Object ref = version_;
      if (ref instanceof java.lang.String) {
        com.google.protobuf.ByteString b =
            com.google.protobuf.ByteString.copyFromUtf8(
                (java.lang.String) ref);
        version_ = b;
        return b;
      } else {
        return (com.google.protobuf.ByteString) ref;
      }
    }

    public static final int PREFER_TOOLBOX_FIELD_NUMBER = 3;
    private boolean preferToolbox_ = false;
    /**
     * <pre>
     * prefer_toolbox indicates whether the editor should be launched with the
     * JetBrains Toolbox instead of JetBrains Gateway
     * </pre>
     *
     * <code>bool prefer_toolbox = 3 [json_name = "preferToolbox"];</code>
     * @return The preferToolbox.
     */
    @java.lang.Override
    public boolean getPreferToolbox() {
      return preferToolbox_;
    }

    private byte memoizedIsInitialized = -1;
    @java.lang.Override
    public final boolean isInitialized() {
      byte isInitialized = memoizedIsInitialized;
      if (isInitialized == 1) return true;
      if (isInitialized == 0) return false;

      memoizedIsInitialized = 1;
      return true;
    }

    @java.lang.Override
    public void writeTo(com.google.protobuf.CodedOutputStream output)
                        throws java.io.IOException {
      if (!com.google.protobuf.GeneratedMessage.isStringEmpty(name_)) {
        com.google.protobuf.GeneratedMessage.writeString(output, 1, name_);
      }
      if (!com.google.protobuf.GeneratedMessage.isStringEmpty(version_)) {
        com.google.protobuf.GeneratedMessage.writeString(output, 2, version_);
      }
      if (preferToolbox_ != false) {
        output.writeBool(3, preferToolbox_);
      }
      getUnknownFields().writeTo(output);
    }

    @java.lang.Override
    public int getSerializedSize() {
      int size = memoizedSize;
      if (size != -1) return size;

      size = 0;
      if (!com.google.protobuf.GeneratedMessage.isStringEmpty(name_)) {
        size += com.google.protobuf.GeneratedMessage.computeStringSize(1, name_);
      }
      if (!com.google.protobuf.GeneratedMessage.isStringEmpty(version_)) {
        size += com.google.protobuf.GeneratedMessage.computeStringSize(2, version_);
      }
      if (preferToolbox_ != false) {
        size += com.google.protobuf.CodedOutputStream
          .computeBoolSize(3, preferToolbox_);
      }
      size += getUnknownFields().getSerializedSize();
      memoizedSize = size;
      return size;
    }

    @java.lang.Override
    public boolean equals(final java.lang.Object obj) {
      if (obj == this) {
       return true;
      }
      if (!(obj instanceof io.nxpod.publicapi.v1.Editor.EditorReference)) {
        return super.equals(obj);
      }
      io.nxpod.publicapi.v1.Editor.EditorReference other = (io.nxpod.publicapi.v1.Editor.EditorReference) obj;

      if (!getName()
          .equals(other.getName())) return false;
      if (!getVersion()
          .equals(other.getVersion())) return false;
      if (getPreferToolbox()
          != other.getPreferToolbox()) return false;
      if (!getUnknownFields().equals(other.getUnknownFields())) return false;
      return true;
    }

    @java.lang.Override
    public int hashCode() {
      if (memoizedHashCode != 0) {
        return memoizedHashCode;
      }
      int hash = 41;
      hash = (19 * hash) + getDescriptor().hashCode();
      hash = (37 * hash) + NAME_FIELD_NUMBER;
      hash = (53 * hash) + getName().hashCode();
      hash = (37 * hash) + VERSION_FIELD_NUMBER;
      hash = (53 * hash) + getVersion().hashCode();
      hash = (37 * hash) + PREFER_TOOLBOX_FIELD_NUMBER;
      hash = (53 * hash) + com.google.protobuf.Internal.hashBoolean(
          getPreferToolbox());
      hash = (29 * hash) + getUnknownFields().hashCode();
      memoizedHashCode = hash;
      return hash;
    }

    public static io.nxpod.publicapi.v1.Editor.EditorReference parseFrom(
        java.nio.ByteBuffer data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static io.nxpod.publicapi.v1.Editor.EditorReference parseFrom(
        java.nio.ByteBuffer data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static io.nxpod.publicapi.v1.Editor.EditorReference parseFrom(
        com.google.protobuf.ByteString data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static io.nxpod.publicapi.v1.Editor.EditorReference parseFrom(
        com.google.protobuf.ByteString data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static io.nxpod.publicapi.v1.Editor.EditorReference parseFrom(byte[] data)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data);
    }
    public static io.nxpod.publicapi.v1.Editor.EditorReference parseFrom(
        byte[] data,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws com.google.protobuf.InvalidProtocolBufferException {
      return PARSER.parseFrom(data, extensionRegistry);
    }
    public static io.nxpod.publicapi.v1.Editor.EditorReference parseFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessage
          .parseWithIOException(PARSER, input);
    }
    public static io.nxpod.publicapi.v1.Editor.EditorReference parseFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessage
          .parseWithIOException(PARSER, input, extensionRegistry);
    }

    public static io.nxpod.publicapi.v1.Editor.EditorReference parseDelimitedFrom(java.io.InputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessage
          .parseDelimitedWithIOException(PARSER, input);
    }

    public static io.nxpod.publicapi.v1.Editor.EditorReference parseDelimitedFrom(
        java.io.InputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessage
          .parseDelimitedWithIOException(PARSER, input, extensionRegistry);
    }
    public static io.nxpod.publicapi.v1.Editor.EditorReference parseFrom(
        com.google.protobuf.CodedInputStream input)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessage
          .parseWithIOException(PARSER, input);
    }
    public static io.nxpod.publicapi.v1.Editor.EditorReference parseFrom(
        com.google.protobuf.CodedInputStream input,
        com.google.protobuf.ExtensionRegistryLite extensionRegistry)
        throws java.io.IOException {
      return com.google.protobuf.GeneratedMessage
          .parseWithIOException(PARSER, input, extensionRegistry);
    }

    @java.lang.Override
    public Builder newBuilderForType() { return newBuilder(); }
    public static Builder newBuilder() {
      return DEFAULT_INSTANCE.toBuilder();
    }
    public static Builder newBuilder(io.nxpod.publicapi.v1.Editor.EditorReference prototype) {
      return DEFAULT_INSTANCE.toBuilder().mergeFrom(prototype);
    }
    @java.lang.Override
    public Builder toBuilder() {
      return this == DEFAULT_INSTANCE
          ? new Builder() : new Builder().mergeFrom(this);
    }

    @java.lang.Override
    protected Builder newBuilderForType(
        com.google.protobuf.GeneratedMessage.BuilderParent parent) {
      Builder builder = new Builder(parent);
      return builder;
    }
    /**
     * Protobuf type {@code nxpod.v1.EditorReference}
     */
    public static final class Builder extends
        com.google.protobuf.GeneratedMessage.Builder<Builder> implements
        // @@protoc_insertion_point(builder_implements:nxpod.v1.EditorReference)
        io.nxpod.publicapi.v1.Editor.EditorReferenceOrBuilder {
      public static final com.google.protobuf.Descriptors.Descriptor
          getDescriptor() {
        return io.nxpod.publicapi.v1.Editor.internal_static_nxpod_v1_EditorReference_descriptor;
      }

      @java.lang.Override
      protected com.google.protobuf.GeneratedMessage.FieldAccessorTable
          internalGetFieldAccessorTable() {
        return io.nxpod.publicapi.v1.Editor.internal_static_nxpod_v1_EditorReference_fieldAccessorTable
            .ensureFieldAccessorsInitialized(
                io.nxpod.publicapi.v1.Editor.EditorReference.class, io.nxpod.publicapi.v1.Editor.EditorReference.Builder.class);
      }

      // Construct using io.nxpod.publicapi.v1.Editor.EditorReference.newBuilder()
      private Builder() {

      }

      private Builder(
          com.google.protobuf.GeneratedMessage.BuilderParent parent) {
        super(parent);

      }
      @java.lang.Override
      public Builder clear() {
        super.clear();
        bitField0_ = 0;
        name_ = "";
        version_ = "";
        preferToolbox_ = false;
        return this;
      }

      @java.lang.Override
      public com.google.protobuf.Descriptors.Descriptor
          getDescriptorForType() {
        return io.nxpod.publicapi.v1.Editor.internal_static_nxpod_v1_EditorReference_descriptor;
      }

      @java.lang.Override
      public io.nxpod.publicapi.v1.Editor.EditorReference getDefaultInstanceForType() {
        return io.nxpod.publicapi.v1.Editor.EditorReference.getDefaultInstance();
      }

      @java.lang.Override
      public io.nxpod.publicapi.v1.Editor.EditorReference build() {
        io.nxpod.publicapi.v1.Editor.EditorReference result = buildPartial();
        if (!result.isInitialized()) {
          throw newUninitializedMessageException(result);
        }
        return result;
      }

      @java.lang.Override
      public io.nxpod.publicapi.v1.Editor.EditorReference buildPartial() {
        io.nxpod.publicapi.v1.Editor.EditorReference result = new io.nxpod.publicapi.v1.Editor.EditorReference(this);
        if (bitField0_ != 0) { buildPartial0(result); }
        onBuilt();
        return result;
      }

      private void buildPartial0(io.nxpod.publicapi.v1.Editor.EditorReference result) {
        int from_bitField0_ = bitField0_;
        if (((from_bitField0_ & 0x00000001) != 0)) {
          result.name_ = name_;
        }
        if (((from_bitField0_ & 0x00000002) != 0)) {
          result.version_ = version_;
        }
        if (((from_bitField0_ & 0x00000004) != 0)) {
          result.preferToolbox_ = preferToolbox_;
        }
      }

      @java.lang.Override
      public Builder mergeFrom(com.google.protobuf.Message other) {
        if (other instanceof io.nxpod.publicapi.v1.Editor.EditorReference) {
          return mergeFrom((io.nxpod.publicapi.v1.Editor.EditorReference)other);
        } else {
          super.mergeFrom(other);
          return this;
        }
      }

      public Builder mergeFrom(io.nxpod.publicapi.v1.Editor.EditorReference other) {
        if (other == io.nxpod.publicapi.v1.Editor.EditorReference.getDefaultInstance()) return this;
        if (!other.getName().isEmpty()) {
          name_ = other.name_;
          bitField0_ |= 0x00000001;
          onChanged();
        }
        if (!other.getVersion().isEmpty()) {
          version_ = other.version_;
          bitField0_ |= 0x00000002;
          onChanged();
        }
        if (other.getPreferToolbox() != false) {
          setPreferToolbox(other.getPreferToolbox());
        }
        this.mergeUnknownFields(other.getUnknownFields());
        onChanged();
        return this;
      }

      @java.lang.Override
      public final boolean isInitialized() {
        return true;
      }

      @java.lang.Override
      public Builder mergeFrom(
          com.google.protobuf.CodedInputStream input,
          com.google.protobuf.ExtensionRegistryLite extensionRegistry)
          throws java.io.IOException {
        if (extensionRegistry == null) {
          throw new java.lang.NullPointerException();
        }
        try {
          boolean done = false;
          while (!done) {
            int tag = input.readTag();
            switch (tag) {
              case 0:
                done = true;
                break;
              case 10: {
                name_ = input.readStringRequireUtf8();
                bitField0_ |= 0x00000001;
                break;
              } // case 10
              case 18: {
                version_ = input.readStringRequireUtf8();
                bitField0_ |= 0x00000002;
                break;
              } // case 18
              case 24: {
                preferToolbox_ = input.readBool();
                bitField0_ |= 0x00000004;
                break;
              } // case 24
              default: {
                if (!super.parseUnknownField(input, extensionRegistry, tag)) {
                  done = true; // was an endgroup tag
                }
                break;
              } // default:
            } // switch (tag)
          } // while (!done)
        } catch (com.google.protobuf.InvalidProtocolBufferException e) {
          throw e.unwrapIOException();
        } finally {
          onChanged();
        } // finally
        return this;
      }
      private int bitField0_;

      private java.lang.Object name_ = "";
      /**
       * <code>string name = 1 [json_name = "name"];</code>
       * @return The name.
       */
      public java.lang.String getName() {
        java.lang.Object ref = name_;
        if (!(ref instanceof java.lang.String)) {
          com.google.protobuf.ByteString bs =
              (com.google.protobuf.ByteString) ref;
          java.lang.String s = bs.toStringUtf8();
          name_ = s;
          return s;
        } else {
          return (java.lang.String) ref;
        }
      }
      /**
       * <code>string name = 1 [json_name = "name"];</code>
       * @return The bytes for name.
       */
      public com.google.protobuf.ByteString
          getNameBytes() {
        java.lang.Object ref = name_;
        if (ref instanceof String) {
          com.google.protobuf.ByteString b =
              com.google.protobuf.ByteString.copyFromUtf8(
                  (java.lang.String) ref);
          name_ = b;
          return b;
        } else {
          return (com.google.protobuf.ByteString) ref;
        }
      }
      /**
       * <code>string name = 1 [json_name = "name"];</code>
       * @param value The name to set.
       * @return This builder for chaining.
       */
      public Builder setName(
          java.lang.String value) {
        if (value == null) { throw new NullPointerException(); }
        name_ = value;
        bitField0_ |= 0x00000001;
        onChanged();
        return this;
      }
      /**
       * <code>string name = 1 [json_name = "name"];</code>
       * @return This builder for chaining.
       */
      public Builder clearName() {
        name_ = getDefaultInstance().getName();
        bitField0_ = (bitField0_ & ~0x00000001);
        onChanged();
        return this;
      }
      /**
       * <code>string name = 1 [json_name = "name"];</code>
       * @param value The bytes for name to set.
       * @return This builder for chaining.
       */
      public Builder setNameBytes(
          com.google.protobuf.ByteString value) {
        if (value == null) { throw new NullPointerException(); }
        checkByteStringIsUtf8(value);
        name_ = value;
        bitField0_ |= 0x00000001;
        onChanged();
        return this;
      }

      private java.lang.Object version_ = "";
      /**
       * <code>string version = 2 [json_name = "version"];</code>
       * @return The version.
       */
      public java.lang.String getVersion() {
        java.lang.Object ref = version_;
        if (!(ref instanceof java.lang.String)) {
          com.google.protobuf.ByteString bs =
              (com.google.protobuf.ByteString) ref;
          java.lang.String s = bs.toStringUtf8();
          version_ = s;
          return s;
        } else {
          return (java.lang.String) ref;
        }
      }
      /**
       * <code>string version = 2 [json_name = "version"];</code>
       * @return The bytes for version.
       */
      public com.google.protobuf.ByteString
          getVersionBytes() {
        java.lang.Object ref = version_;
        if (ref instanceof String) {
          com.google.protobuf.ByteString b =
              com.google.protobuf.ByteString.copyFromUtf8(
                  (java.lang.String) ref);
          version_ = b;
          return b;
        } else {
          return (com.google.protobuf.ByteString) ref;
        }
      }
      /**
       * <code>string version = 2 [json_name = "version"];</code>
       * @param value The version to set.
       * @return This builder for chaining.
       */
      public Builder setVersion(
          java.lang.String value) {
        if (value == null) { throw new NullPointerException(); }
        version_ = value;
        bitField0_ |= 0x00000002;
        onChanged();
        return this;
      }
      /**
       * <code>string version = 2 [json_name = "version"];</code>
       * @return This builder for chaining.
       */
      public Builder clearVersion() {
        version_ = getDefaultInstance().getVersion();
        bitField0_ = (bitField0_ & ~0x00000002);
        onChanged();
        return this;
      }
      /**
       * <code>string version = 2 [json_name = "version"];</code>
       * @param value The bytes for version to set.
       * @return This builder for chaining.
       */
      public Builder setVersionBytes(
          com.google.protobuf.ByteString value) {
        if (value == null) { throw new NullPointerException(); }
        checkByteStringIsUtf8(value);
        version_ = value;
        bitField0_ |= 0x00000002;
        onChanged();
        return this;
      }

      private boolean preferToolbox_ ;
      /**
       * <pre>
       * prefer_toolbox indicates whether the editor should be launched with the
       * JetBrains Toolbox instead of JetBrains Gateway
       * </pre>
       *
       * <code>bool prefer_toolbox = 3 [json_name = "preferToolbox"];</code>
       * @return The preferToolbox.
       */
      @java.lang.Override
      public boolean getPreferToolbox() {
        return preferToolbox_;
      }
      /**
       * <pre>
       * prefer_toolbox indicates whether the editor should be launched with the
       * JetBrains Toolbox instead of JetBrains Gateway
       * </pre>
       *
       * <code>bool prefer_toolbox = 3 [json_name = "preferToolbox"];</code>
       * @param value The preferToolbox to set.
       * @return This builder for chaining.
       */
      public Builder setPreferToolbox(boolean value) {

        preferToolbox_ = value;
        bitField0_ |= 0x00000004;
        onChanged();
        return this;
      }
      /**
       * <pre>
       * prefer_toolbox indicates whether the editor should be launched with the
       * JetBrains Toolbox instead of JetBrains Gateway
       * </pre>
       *
       * <code>bool prefer_toolbox = 3 [json_name = "preferToolbox"];</code>
       * @return This builder for chaining.
       */
      public Builder clearPreferToolbox() {
        bitField0_ = (bitField0_ & ~0x00000004);
        preferToolbox_ = false;
        onChanged();
        return this;
      }

      // @@protoc_insertion_point(builder_scope:nxpod.v1.EditorReference)
    }

    // @@protoc_insertion_point(class_scope:nxpod.v1.EditorReference)
    private static final io.nxpod.publicapi.v1.Editor.EditorReference DEFAULT_INSTANCE;
    static {
      DEFAULT_INSTANCE = new io.nxpod.publicapi.v1.Editor.EditorReference();
    }

    public static io.nxpod.publicapi.v1.Editor.EditorReference getDefaultInstance() {
      return DEFAULT_INSTANCE;
    }

    private static final com.google.protobuf.Parser<EditorReference>
        PARSER = new com.google.protobuf.AbstractParser<EditorReference>() {
      @java.lang.Override
      public EditorReference parsePartialFrom(
          com.google.protobuf.CodedInputStream input,
          com.google.protobuf.ExtensionRegistryLite extensionRegistry)
          throws com.google.protobuf.InvalidProtocolBufferException {
        Builder builder = newBuilder();
        try {
          builder.mergeFrom(input, extensionRegistry);
        } catch (com.google.protobuf.InvalidProtocolBufferException e) {
          throw e.setUnfinishedMessage(builder.buildPartial());
        } catch (com.google.protobuf.UninitializedMessageException e) {
          throw e.asInvalidProtocolBufferException().setUnfinishedMessage(builder.buildPartial());
        } catch (java.io.IOException e) {
          throw new com.google.protobuf.InvalidProtocolBufferException(e)
              .setUnfinishedMessage(builder.buildPartial());
        }
        return builder.buildPartial();
      }
    };

    public static com.google.protobuf.Parser<EditorReference> parser() {
      return PARSER;
    }

    @java.lang.Override
    public com.google.protobuf.Parser<EditorReference> getParserForType() {
      return PARSER;
    }

    @java.lang.Override
    public io.nxpod.publicapi.v1.Editor.EditorReference getDefaultInstanceForType() {
      return DEFAULT_INSTANCE;
    }

  }

  private static final com.google.protobuf.Descriptors.Descriptor
    internal_static_nxpod_v1_EditorReference_descriptor;
  private static final
    com.google.protobuf.GeneratedMessage.FieldAccessorTable
      internal_static_nxpod_v1_EditorReference_fieldAccessorTable;

  public static com.google.protobuf.Descriptors.FileDescriptor
      getDescriptor() {
    return descriptor;
  }
  private static  com.google.protobuf.Descriptors.FileDescriptor
      descriptor;
  static {
    java.lang.String[] descriptorData = {
      "\n\026nxpod/v1/editor.proto\022\tnxpod.v1\"f\n\017E" +
      "ditorReference\022\022\n\004name\030\001 \001(\tR\004name\022\030\n\007ve" +
      "rsion\030\002 \001(\tR\007version\022%\n\016prefer_toolbox\030\003" +
      " \001(\010R\rpreferToolboxBQ\n\026io.nxpod.publica" +
      "pi.v1Z7github.com/nxpkg/nxpod/compo" +
      "nents/public-api/go/v1b\006proto3"
    };
    descriptor = com.google.protobuf.Descriptors.FileDescriptor
      .internalBuildGeneratedFileFrom(descriptorData,
        new com.google.protobuf.Descriptors.FileDescriptor[] {
        });
    internal_static_nxpod_v1_EditorReference_descriptor =
      getDescriptor().getMessageTypes().get(0);
    internal_static_nxpod_v1_EditorReference_fieldAccessorTable = new
      com.google.protobuf.GeneratedMessage.FieldAccessorTable(
        internal_static_nxpod_v1_EditorReference_descriptor,
        new java.lang.String[] { "Name", "Version", "PreferToolbox", });
    descriptor.resolveAllFeaturesImmutable();
  }

  // @@protoc_insertion_point(outer_class_scope)
}
